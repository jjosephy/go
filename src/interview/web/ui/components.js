// Header Component //
var Header = React.createClass({displayName: 'Header',
    newHandler : function(e) {
        this.setDialog('new', 'Create New Interview', 'Save');
    },
    editHandler : function(e) {
        this.setDialog('edit', 'Edit Interview', 'Save');
    },
    findHandler : function(e) {
        this.setDialog('find', 'Find Interview', 'Go');
    },
    setDialog : function(type, title, buttonText) {
        var e = this.props.dialog;
        if (e) {
            e.setState({ contentType: type });
            e.setState({ showDialog: true });
            e.setState({ title: title });
            e.setState({ buttonText: buttonText });
        } else {
            console.log('no dialog object');
        }
    },
    render: function() {
        return (
            <div className="header" >
                <MenuItem label="New" handler={this.newHandler} />
                <MenuItem label="Edit" handler={this.editHandler} />
                <MenuItem label="Find" handler={this.findHandler} />
            </div>
        );
    }
});

// New Form //
var NewForm = React.createClass({
    render: function() {
        return (
            <div>
                <div>Candidate Name</div>
                <input type="text" className="iText" ref="cname"/>
                <div className="interviewers">
                    <span className="lInterview">Interviewers</span>
                </div>
                <div>
                    &nbsp;
                </div>
                <div>
                    <div className="addlabel">
                        <label>Interviewer One</label>
                        <input type="text" className="iText" ref="ic1" />
                    </div>
                    <div className="addlabel">
                        <label>Interviewer Two</label>
                        <input type="text" className="iText" ref="ic2" />
                    </div>
                    <div className="addlabel">
                        <label>Interviewer Three</label>
                        <input type="text" className="iText" ref="ic3" />
                    </div>
                    <div className="addlabel">
                        <label>Interviewer Four</label>
                        <input type="text" className="iText" ref="ic4" />
                    </div>
                    <div className="addlabel">
                        <label>Interviewer Five</label>
                        <input type="text" className="iText" ref="ic5" />
                    </div>
                </div>
            </div>
        );
    }
});

// Edit Form //
var EditForm = React.createClass({
    render: function() {
        return (
            <div>
                Edit
                <input type="text"/>
            </div>
        );
    }
});

// Find Form//
var FindForm = React.createClass({
    render: function() {
        return (
            <div>
                <div className="addlabel">
                    <label>Interviewer Id</label>
                    <input type="text" className="iText" ref="ic1" />
                </div>
                <div className="addlabel">
                    <label>Candidate Name</label>
                    <input type="text" className="iText" ref="ic2" />
                </div>
            </div>
        );
    }
});

// Dialog //
var Dialog = React.createClass({
    getInitialState: function() {
        return {
            showDialog: false,
            contentType: 'none',
            title: 'title',
            buttonText: 'Save',
        };
    },
    saveInterview : function() {
        var cname = this.refs.body.refs.cname.getDOMNode().value;
        if (!cname || cname.trim().length == 0) {
            alert('Must provide candidate name');
            return false;
        }

        var i = {
            candidate: cname,
            complete: false,
            comments: new Array(),
        };

        var isEmpty = true;
        for (var x = 1; x < 6; x++) {
            var z = 'ic' + x.toString();
            var n = this.refs.body.refs[z].getDOMNode().value;
            if (n && n.length > 0) {
                isEmpty = false;
                var c = {
                    content: '',
                    interviewer: n
                };
                i.comments.push(c);
            }
        }

        if (isEmpty) {
            alert('Need to supply at least one interviewer.')
            return false;
        }

        var that = this;
        Client.SaveInterview(
            i,
            function(data, textStatus, jqXHR) {
                var res = JSON.parse(jqXHR.responseText);
                React.render(
                    <div>
                        <div> Candidate Name: {res.candidate}</div>
                        {
                            res.comments.map(function(c) {
                                return <CommentPanel body={c}></CommentPanel>
                            })
                        }
                    </div>,
                    document.getElementById('content')
                );
                document.getElementById('footer').innerText = 'Interview ' + res.id + ' saved successfully';
            },
            function (jqXHR, textStatus, errorThrown) {
                var res = JSON.parse(jqXHR.responseText);
                var msg = "Error Code: " + res.ErrorCode + " Message: " + res.Message;
                console.log(msg);
                document.getElementById('footer').innerText = msg;
                that.showErrorContent(res);
            }
        );
        this.showLoadingPanel();
        return true;
    },
    showLoadingPanel() {
        React.render(
            <LoadingPanel />,
            document.getElementById('content')
        );
    },
    showErrorContent(res) {
        React.render(
            <div>
                {res.Message}
            </div>,
            document.getElementById('content')
        );
    },
    getInterview() {
        var id = this.refs.body.refs.ic1.getDOMNode().value;
        var that = this;
        Client.GetInterview(
            id,
            '',
            function(data, textStatus, jqXHR) {
                var res = JSON.parse(jqXHR.responseText);
                React.render(
                    <div>
                        <div> Candidate Name: {res.candidate}</div>
                        {
                            res.comments.map(function(c) {
                                return <CommentPanel body={c}></CommentPanel>
                            })
                        }
                    </div>,
                    document.getElementById('content')
                );
                document.getElementById('footer').innerText = 'Interview ' + id + ' retrieved successfully';
            },
            function (jqXHR, textStatus, errorThrown) {
                var res = JSON.parse(jqXHR.responseText);
                var msg = "Error Code: " + res.ErrorCode + " Message: " + res.Message;
                console.log(msg);
                document.getElementById('footer').innerText = msg;
                that.showErrorContent(res);
            }
        );
        this.showLoadingPanel();
    },
    handleSave : function(e) {
        var success = false;
        switch (this.state.contentType) {
            case 'new':
                success = this.saveInterview();
                break;
            case 'edit':
            case 'find':
                this.getInterview();
                success = true;
                break;
        }

        if (success === true) {
            this.setState({ showDialog: false });
        }
    },
    handleCancel : function(e) {
        switch (this.state.contentType) {
            case 'new':
                for (var x = 1; x < 6; x++) {
                    var i = 'ic' + x.toString();
                    this.refs.body.refs[i].getDOMNode().value = '';
                }
            case 'edit':
            case 'find':
        }

        this.setState({ showDialog: false });
    },
    render: function() {
        return (
            <div>
                { this.state.showDialog ?
                    <div className="parentDialog">
                        <div className="newDialog">
                            <div className="dialogHeader">
                                {this.state.title}
                            </div>
                            <div className="dialogBody">
                                {(() => {
                                    switch (this.state.contentType) {
                                        case 'none':    return 'none';
                                        case 'new':     return <NewForm ref="body"/>;
                                        case 'edit':    return <EditForm ref="body"/>;
                                        case 'find':    return <FindForm ref="body"/>
                                        default:        return 'empty';
                                    }
                                })()}
                            </div>
                            <div className="dialogFooter">
                                <button onClick={this.handleSave} className="dialogButton">{this.state.buttonText}</button>
                                <button onClick={this.handleCancel} className="dialogButton">Cancel</button>
                            </div>
                        </div>
                    </div>
                : null }
            </div>
        );
    }
});

// Menu Item Component //
var MenuItem  = React.createClass({
    handleClick : function(e) {
        this.props.handler();
    },
    render: function() {
        return (
            <span onClick={this.handleClick} className="menu_item">{this.props.label}</span>
        );
    }
});

// Component Panel //
var CommentPanel = React.createClass({
    render: function() {
        return (
            <div className="commentPanel">
                <div className="cpLabel">Interviewer Name: {this.props.body.interviewer}</div>
                <div>
                    <textarea className="txt">
                    </textarea>
                </div>
                <div className="select"><select>
                        <option>No Hire</option>
                        <option>Hire</option>
                    </select>

                    <button className="buttonRight">Update Comments</button>
                </div>
                <hr/>
            </div>
        );
    }
});

// Loading Panel //
var LoadingPanel = React.createClass({
    render: function() {
        return (
            <div className="loading">
                <img src="images/loading.gif" />             
            </div>
        );
    }
});

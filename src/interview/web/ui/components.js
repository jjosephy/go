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
                New
                <input type="text"/>
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
                <label className="dialogLabel">Id</label>
                <input type="text" className="textBox"/>
                <label className="dialogLabel">Candidate Name</label>
                <input type="text" className="textBox"/>
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
            buttonText: 'Save'
        };
    },
    handleClick : function(e) {
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
                                        case 'new':     return <NewForm />;
                                        case 'edit':    return <EditForm />;
                                        case 'find':    return <FindForm />
                                        default:        return 'empty';
                                    }
                                })()}
                            </div>
                            <div className="dialogFooter">
                                <button onClick={this.handleClick} className="dialogButton">{this.state.buttonText}</button>
                                <button onClick={this.handleClick} className="dialogButton">Cancel</button>
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
                <textarea>
                </textarea>
            </div>
        );
    }
});

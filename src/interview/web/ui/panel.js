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

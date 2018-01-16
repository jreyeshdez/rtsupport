import React, {Component} from 'react';
import PropTypes from 'prop-types';
import fecha from 'fecha';

class Message extends Component {
    render() {
        let {message} = this.props;
        fecha.masks.myMask = 'MM-DD-YYYY HH:mm:ss A';
        let date = fecha.parse(message.createdat, 'myMask');
        let createdAt = fecha.format(date, 'HH:mm:ss MM/DD/YY');
        return (
            <li className='message'>
                <div className='author'>
                    <strong>{message.author}</strong>
                    <i className='timestamp'>{createdAt}</i>
                </div>
                <div className='body'>{message.body}</div>
            </li>
        )
    }
}

Message.propTypes = {
    message: PropTypes.object.isRequired
}

export default Message


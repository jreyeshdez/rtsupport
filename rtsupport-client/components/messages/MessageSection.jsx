import React, {Component} from 'react';
import PropTypes from 'prop-types';
import MessageForm from './MessageForm.jsx';
import MessageList from './MessageList.jsx';

class MessageSection extends Component {
    render() {
        let {activeChannel} = this.props;
        return (
            <div className='messages-container panel-default'>
                <div className='panel-heading'>
                    <strong>{activeChannel.name || 'Select a Channel'}</strong>
                </div>
                <div className='panel-body messages'>
                    <MessageList {...this.props} />
                    <MessageForm {...this.props} />
                </div>
            </div>
        )
    }
}

MessageSection.propTypes = {
    messages: PropTypes.array.isRequired,
    addMessage: PropTypes.func.isRequired,
    activeChannel: PropTypes.object.isRequired
}

export default MessageSection
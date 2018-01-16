import React, {Component} from 'react';
import PropTypes from 'prop-types';

class MessageForm extends Component {
    onSubmit(e) {
        e.preventDefault();
        const node = this.refs.message;
        const message = node.value;
        this.props.addMessage(message);
        node.value = '';
    }

    render() {
        let input;
        if(this.props.activeChannel.id !== undefined) {
            input = (
                <input 
                    className='form-control' 
                    placeholder='Whats on your mind...?' 
                    type='text' 
                    ref='message' />    
            )
        }
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
                <div className='form-group'>{input}</div>
            </form>
        )
    }
}

MessageForm.propTypes = {
    addMessage: PropTypes.func.isRequired,
    activeChannel: PropTypes.object.isRequired
}

export default MessageForm
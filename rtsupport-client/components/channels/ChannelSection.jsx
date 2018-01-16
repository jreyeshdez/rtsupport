import React, {Component} from 'react';
import PropTypes from 'prop-types';
import ChannelForm from './ChannelForm.jsx';
import ChannelList from './ChannelList.jsx';

class ChannelSection extends Component {
    onSubmit(e) {
        e.preventDefault();
        const node = this.refs.channel;
        const channelName = node.value;
        this.props.addChannel(channelName);
        node.value = ''
    }

    render() {
        return (
            <div className='support panel panel-primary'>
                <div className='panel-heading'>
                    <strong>Channels</strong>
                </div>
                <div className='panel-body channels'>
                    <ChannelList {...this.props} />
                    <ChannelForm {...this.props} />
                </div>
            </div>
        )
    }
}

ChannelSection.propTypes = {
    channels: PropTypes.array.isRequired,
    setChannel: PropTypes.func.isRequired,
    addChannel: PropTypes.func.isRequired,
    activeChannel: PropTypes.object.isRequired
}

export default ChannelSection
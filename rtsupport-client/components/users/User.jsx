import React, {Component} from 'react';
import PropTypes from 'prop-types';

class User extends Component {

    render() {
        const {user} = this.props;
        return (
            <li>
                {user.name}
            </li>
        )
    }
}

User.propTypes = {
    user: PropTypes.object.isRequired
}

export default User


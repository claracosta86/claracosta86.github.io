import React from 'react';
import { Link } from 'react-router-dom';
import logo from '../../assets/logo.png';
import logoutIcon from '../../assets/logout-icon.png';
import notificationsIcon from '../../assets/notifications-icon.png';

const Header = ({ onLogoutClick, onNotificationClick }) => {
  return (
    <header className="top-bar">
      <img src={logo} alt="Logo Cultural" className="logo-tiny" />
      <div className="right-section">
        <div onClick={onLogoutClick} className="icon-button-container">
          <img src={logoutIcon} alt="Log-out" className="icon" />
        </div>
        <div onClick={onNotificationClick} className="icon-button-container">
          <img
            src={notificationsIcon}
            id="notifications-icon"
            alt="Notificações"
            className="icon"
          />
        </div>
      </div>
    </header>
  );
};

export default Header;

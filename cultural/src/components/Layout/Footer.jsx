import React from 'react';
import { Link } from 'react-router-dom';
import homeIcon from '../../assets/home-icon.png';
import searchIcon from '../../assets/search-icon.png';
import addIcon from '../../assets/add-icon.png';
import favoriteIcon from '../../assets/favorite-icon.png';
import userIcon from '../../assets/user-icon.png';

const Footer = ({ userType }) => {
  return (
    <footer className="footer">
      <Link to={`/home`}>
        <img src={homeIcon} alt="Logo Cultural" />
      </Link>
      <Link to={`/search`}>
        <img src={searchIcon} alt="Buscar" />
      </Link>
      {userType === 'organizer' && (
        <Link to={`/create-cultural`}>
          <img src={addIcon} alt="Adicionar" className="mostImportantButton" />
        </Link>
      )}
      <Link to={`/user/favorites`}>
        <img src={favoriteIcon} alt="Favoritos" />
      </Link>
      <Link to={`/user/profile`}>
        <img src={userIcon} alt="Usuário" />
      </Link>
    </footer>
  );
};

export default Footer;

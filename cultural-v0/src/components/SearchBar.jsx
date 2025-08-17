import { useState } from 'react';

const SearchBar = ({ onSearch }) => {
  const [searchTerm, setSearchTerm] = useState('');

  const handleSearch = (e) => {
    e.preventDefault();
    if (onSearch) {
      onSearch(searchTerm);
    }
  };

  const handleInputChange = (e) => {
    setSearchTerm(e.target.value);
    // Busca em tempo real se necessário
    if (onSearch) {
      onSearch(e.target.value);
    }
  };

  return (
    <section className="search-bar">
      <img src="/images/search-icon.png" alt="Buscar" className="search-icon" />
      <form onSubmit={handleSearch}>
        <input
          id="search-input"
          type="text"
          placeholder="Buscar..."
          value={searchTerm}
          onChange={handleInputChange}
        />
      </form>
    </section>
  );
};

export default SearchBar;


import { useState, useCallback, useMemo } from 'react';

export const useSearch = (data = []) => {
  const [searchTerm, setSearchTerm] = useState('');
  const [searchFilters, setSearchFilters] = useState({
    type: 'all', // 'all', 'event', 'attraction'
    category: 'all',
    date: 'all'
  });

  const handleSearch = useCallback((term) => {
    setSearchTerm(term);
  }, []);

  const handleFilterChange = useCallback((filterType, value) => {
    setSearchFilters(prev => ({
      ...prev,
      [filterType]: value
    }));
  }, []);

  const clearSearch = useCallback(() => {
    setSearchTerm('');
    setSearchFilters({
      type: 'all',
      category: 'all',
      date: 'all'
    });
  }, []);

  const filteredData = useMemo(() => {
    if (!data || data.length === 0) return [];

    let filtered = data;

    // Filtrar por termo de busca
    if (searchTerm.trim()) {
      filtered = filtered.filter(item => 
        item.title?.toLowerCase().includes(searchTerm.toLowerCase()) ||
        item.description?.toLowerCase().includes(searchTerm.toLowerCase()) ||
        item.category?.toLowerCase().includes(searchTerm.toLowerCase())
      );
    }

    // Filtrar por tipo
    if (searchFilters.type !== 'all') {
      filtered = filtered.filter(item => item.type === searchFilters.type);
    }

    // Filtrar por categoria
    if (searchFilters.category !== 'all') {
      filtered = filtered.filter(item => item.category === searchFilters.category);
    }

    // Filtrar por data (implementação básica)
    if (searchFilters.date !== 'all') {
      const today = new Date();
      const tomorrow = new Date(today);
      tomorrow.setDate(tomorrow.getDate() + 1);
      
      switch (searchFilters.date) {
        case 'today':
          filtered = filtered.filter(item => {
            if (!item.date) return false;
            const itemDate = new Date(item.date);
            return itemDate.toDateString() === today.toDateString();
          });
          break;
        case 'tomorrow':
          filtered = filtered.filter(item => {
            if (!item.date) return false;
            const itemDate = new Date(item.date);
            return itemDate.toDateString() === tomorrow.toDateString();
          });
          break;
        case 'this-week':
          const endOfWeek = new Date(today);
          endOfWeek.setDate(today.getDate() + 7);
          filtered = filtered.filter(item => {
            if (!item.date) return false;
            const itemDate = new Date(item.date);
            return itemDate >= today && itemDate <= endOfWeek;
          });
          break;
        default:
          break;
      }
    }

    return filtered;
  }, [data, searchTerm, searchFilters]);

  const searchStats = useMemo(() => {
    const total = data?.length || 0;
    const found = filteredData.length;
    const hasResults = found > 0;
    const hasFilters = searchTerm.trim() || Object.values(searchFilters).some(v => v !== 'all');

    return {
      total,
      found,
      hasResults,
      hasFilters,
      isEmpty: total === 0,
      isFiltered: hasFilters
    };
  }, [data, filteredData, searchTerm, searchFilters]);

  return {
    searchTerm,
    searchFilters,
    filteredData,
    searchStats,
    handleSearch,
    handleFilterChange,
    clearSearch
  };
};


import React, { useState } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import './styles/create.css'; // Certifique-se que o CSS está correto
import logo from '../assets/logo.png';
import gobackIcon from '../assets/goback.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';

const NotificationModal = ({ isOpen, onClose, notifications }) => {
  if (!isOpen) return null;

  const handleLinkClick = (culturalID) =>  async () => {
    navigate(`/card/${culturalID}`);
  };
  
  return (
     <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={e => e.stopPropagation()}>
        <h2 className="modal-title">
          Notificações
        </h2>

        <div className="modal-content">
          {notifications.length === 0 ? (
            <p>Você não tem novas notificações.</p>
          ) : (
            notifications.map((notif, index) => (
              <div key={index} className="notification-item">
                <p> Veja as atualizações de <button onClick={() => handleLinkClick(notif.ID)}>{notif.Title}</button></p>
              </div>
            ))
          )}
        </div>

        <div className="modal-actions">
          <button
            onClick={onClose}
            className="modal-close-btn"
          >
            Entendi
          </button>
        </div>
      </div>
    </div>
  );
};

const CreateCulturalPage = () => {
  const navigate = useNavigate();
  const location = useLocation();
  
  const { userType, userID } = location.state || {};

  const [selectedFile, setSelectedFile] = useState(null);
  const [culturalType, setCulturalType] = useState('');
  const [error, setError] = useState('');
  
  // Estado único para gerenciar todos os campos do formulário
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    location: '',
    price: 0,
    isAccessible: false,
    startDate: '',
    endDate: '',
    durationTime: '',
    openDays: [],
    openTime: '',
  });

  const handleTypeChange = (event) => {
    setCulturalType(event.target.value);
    // Limpa os dados ao trocar o tipo para evitar enviar dados errados
    setFormData({
      title: '', description: '', location: '', price: 'R$0,00', isAccessible: false,
      startDate: '', endDate: '', durationTime: '',
      openDays: [], openTime: '',
    });
  };

  const handleInputChange = (event) => {
    const { name, value, type, checked } = event.target;
    const finalValue = type === 'checkbox' ? checked : value;
    setFormData(prevData => ({ ...prevData, [name]: finalValue }));
  };

  const handleCheckboxChange = (event) => {
    const { value, checked } = event.target;
    setFormData(prevData => {
      const currentDays = prevData.openDays;
      if (checked) {
        return { ...prevData, openDays: [...currentDays, value] };
      } else {
        return { ...prevData, openDays: currentDays.filter(day => day !== value) };
      }
    });
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError(''); 

    const submissionFormData = new FormData();

    if (selectedFile) {
      submissionFormData.append('image', selectedFile);
    }
    
    let finalPayload = {
      Title: formData.title,
      Type: culturalType,
      Description: formData.description,
      Price: parseFloat(formData.price),
      IsAccessible: formData.isAccessible,
      Organizer: {
        ID: userID
      },
    };

    if (culturalType === 'event') {
      finalPayload.Location = formData.location;
      finalPayload.Event = {
        StartDate: formData.startDate,
        EndDate: formData.endDate,
        DurationTime: formData.durationTime,
      };
    } else if (culturalType === 'tourist_attraction') {
      finalPayload.Location = formData.location; 
      finalPayload.TouristAttraction = {
        OpenDays: formData.openDays.join(', '),
        OpenTime: formData.openTime,
      };
    }
    
    submissionFormData.append('data', JSON.stringify(finalPayload));
    
    console.log('Enviando para a API:', JSON.stringify(finalPayload, null, 2));

    try {
      const response = await fetch('http://localhost:8080/culturais/', { 
        method: 'POST',
        body: submissionFormData,
        credentials: 'include',
      });

      if (response.ok) {
        const result = await response.json();
        alert('Item cultural criado com sucesso!');
        navigate(`/card/${result.ID}`, { state: { userID, userType } });
      } else {
        const errorData = await response.json();
        console.error('Erro da API:', errorData);
        setError(errorData.message || 'Falha ao criar o item.');
        alert(`Falha ao criar o item: ${errorData.message}`);
      }
    } catch (err) {
      console.error('Erro de rede:', err);
      setError('Não foi possível se conectar ao servidor.');
      alert('Não foi possível se conectar ao servidor.');
    }
  };

  const handleUserIconClick = async () => {
    navigate('/user/profile',  { state: {userID: userID, userType: userType} });
  };

  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

  const handleNotificationIconClick = async () => {
    const fetchNewNotifications = async () => {
        try {
          const response = await fetch(`http://localhost:8080/notifications/${userID}`, {
               credentials: 'include'
          });
          if (response.ok) {
            const data = await response.json();
            setNotification(data.culturals);
            console.log("Notificações recebidas:", data.culturals);
          }
        } catch (error) {
          console.error("Erro ao buscar por novas notificações:", error);
        }
      };
      fetchNewNotifications();
    setNotificationModalOpen(true, notification);
  };
  
  const daysOfWeek = ['Segunda', 'Terça', 'Quarta', 'Quinta', 'Sexta', 'Sábado', 'Domingo'];

  return (
    <>
    <NotificationModal isOpen={isNotificationModalOpen} onClose={() => setNotificationModalOpen(false)} notifications={notification} />
    <section className="screen" id="tela-create">
      <header className='top-bar-common'>
          <div className="logo-container">
            <Link to="/home">
              <img src={logo} alt="Logo Cultural" className="logo-tiny" />
            </Link>
          </div>
          <div className="right-section">
            <div className="icons">
              <div onClick={handleNotificationIconClick} className="icon-button-container">
                <img src={notificationsIcon} id="notifications-icon" alt="Notificações" className="icon" />
              </div>
              <div onClick={handleUserIconClick} className="icon-button-container">
                <img src={userIcon} id="user-icon" alt="Usuário" className="icon" />
              </div>
          </div>
        </div>
      </header>

      <div className="create-box">
        <div className="header-title">
          <button onClick={() => navigate(-1)} className="goback-btn">
            <img src={gobackIcon} alt="Voltar" className="goback-img" />
          </button>
          <h2>Crie Seu Cultural</h2>
        </div>
        
        <form id="createCulturalForm" onSubmit={handleSubmit}>
          <fieldset className="selection-box">
            <legend>Selecione seu tipo de cultural:</legend>
             <div className="radio-option">
                <input type="radio" id="event" name="culturalType" value="event" checked={culturalType === 'event'} onChange={handleTypeChange} />
                <label htmlFor="event">Evento</label>
            </div>
            <div className="radio-option">
                <input type="radio" id="tourist_attraction" name="culturalType" value="tourist_attraction" checked={culturalType === 'tourist_attraction'} onChange={handleTypeChange}/>
                <label htmlFor="tourist_attraction">Ponto Turístico</label>
            </div>
          </fieldset>
          
          {culturalType && (
            <>
              <label htmlFor="title" className="required">Nome</label>
              <input id="title" name="title" type="text" placeholder="Nome do evento ou local" value={formData.title} onChange={handleInputChange} required />

              <label htmlFor="location" className="required">Endereço</label>
              <input id="location" name="location" type="text" placeholder="Rua, número, bairro..." value={formData.location} onChange={handleInputChange} required />

              {culturalType === 'event' ? (
                <>
                  <label htmlFor="startDate" className="required">Data e Hora de Início</label>
                  <input type="datetime-local" id="startDate" name="startDate" value={formData.startDate} onChange={handleInputChange} required />

                  <label htmlFor="endDate" className="required">Data e Hora de Fim</label>
                  <input type="datetime-local" id="endDate" name="endDate" value={formData.endDate} onChange={handleInputChange} required />

                  <label htmlFor="durationTime" className="required">Duração</label>
                  <input type="text" id="durationTime" name="durationTime" placeholder="HH:MM" value={formData.durationTime} onChange={handleInputChange} required />
                </>
              ) : (
                <>
                  <label className="required">Dias de Funcionamento</label>
                  <div className="checkbox-group">
                    {daysOfWeek.map(day => (
                      <div key={day} className="checkbox-option">
                        <input type="checkbox" id={day} value={day} checked={formData.openDays.includes(day)} onChange={handleCheckboxChange} />
                        <label htmlFor={day}>{day}</label>
                      </div>
                    ))}
                  </div>

                  <label htmlFor="openTime" className="required">Horário de Funcionamento</label>
                  <input type="text" id="openTime" name="openTime" placeholder="ex: 09:00 às 17:00" value={formData.openTime} onChange={handleInputChange} required />
                </>
              )}

              <label htmlFor="description">Descrição</label>
              <textarea id="description" name="description" placeholder="Descreva seu cultural :)" value={formData.description} onChange={handleInputChange} />

              <label htmlFor="price">Preço</label>
              <input type="text" id="price" name="price" value={formData.price} onChange={handleInputChange} required />
              
              <label htmlFor="accessibility">Acessibilidade</label>
              <div className="checkbox-option accessibility-option">
                <input type="checkbox" id="isAccessible" name="isAccessible" checked={formData.isAccessible} onChange={handleInputChange} />
                <label htmlFor="isAccessible">Possui estrutura de acessibilidade</label>
              </div>

              <label htmlFor="image">Imagem do Cultural</label>
              <input type="file" id="image" name="image" onChange={(event) => setSelectedFile(event.target.files[0])} />

              {error && <span className="error">{error}</span>}
              
              <p className="button-container">
                <button type="submit" className="btn">Criar</button>
              </p>
            </>
          )}
        </form>
      </div>
    </section>
    </>
  );
};

export default CreateCulturalPage;
import { useNavigate } from 'react-router-dom';

const Card = ({ 
  title, 
  image, 
  alt, 
  type = 'event', // 'event' ou 'attraction'
  id,
  onClick 
}) => {
  const navigate = useNavigate();

  const handleClick = () => {
    if (onClick) {
      onClick();
    } else {
      // Navegação padrão baseada no tipo
      if (type === 'event') {
        navigate(`/event/${id}`);
      } else {
        navigate(`/attraction/${id}`);
      }
    }
  };

  return (
    <div className="card" onClick={handleClick}>
      <p className="title">{title}</p>
      <img src={image} alt={alt} />
    </div>
  );
};

export default Card;


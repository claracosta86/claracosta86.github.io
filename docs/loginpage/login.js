// document.addEventListener('DOMContentLoaded', () => {
//     const userTypeRaw = sessionStorage.getItem('userType');

//     // whitelist para evitar valores inesperados
//     const userType = (userTypeRaw === 'organizer' || userTypeRaw === 'common') ? userTypeRaw : 'common';
//     window.location.href = `../register/register-${userType}.html`;
// });

// document.addEventListener('DOMContentLoaded', function () {
//     const loginForm = document.getElementById('loginForm');
//     loginForm.addEventListener('submit', function(event) {
//         event.preventDefault();
   
//         const btn = event.submitter;     
//         const userType = btn?.dataset.role || 'common';
//         sessionStorage.setItem('userType', userType);
//   });
// });

document.addEventListener('DOMContentLoaded', function () {
    const registerForm = document.getElementById('indexForm');
    registerForm.addEventListener('submit', function(event) {
        event.preventDefault();
   
        const btn = event.submitter;     
        const userType = btn?.dataset.role || 'common';
        sessionStorage.setItem('userType', role);            

        window.location.href = `../login/login.html?role=${encodeURIComponent(userType)}`;
  });
});

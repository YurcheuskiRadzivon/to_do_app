document.addEventListener('DOMContentLoaded', function () {
    const form = document.getElementById('loginForm');
    
    form.addEventListener('submit', function (event) {
        event.preventDefault(); 

        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;

        
        if (email === '' || password === '') {
            alert('All fields are required');
            return;
        }
        if (!validateEmail(email)) {
            alert('Incorrect email');
            return;
        }
        if (password.length < 6) {
            alert('Incorrect password');
            return;
        }
        /*if (!validateUsername(username)) {
            alert('Username can only contain letters and numbers');
            return;
        }*/

        const data = {
            email: email,
            password: password
        };

        fetch('/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(response => response.json())
        .then(result => {
            if (result.error) {
                alert(result.error); 
            } else {
                alert("login successful"); 
                window.location.href = '/redirect'; 
            }
        })
        .catch(error => {
            alert(`Error: ${error}`); 
            window.location.href = '/';
        });
    });

    function validateEmail(email) {
        const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return re.test(String(email).toLowerCase());
    }
   /* function validateUsername(username) {
        const re = /^[a-zA-Z0-9а-яА-Я]+$/; 
        return re.test(username);
    }*/
});

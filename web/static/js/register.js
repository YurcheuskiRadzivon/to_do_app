document.addEventListener('DOMContentLoaded', function () {
    const form = document.getElementById('registerForm');

    form.addEventListener('submit', function (event) {
        event.preventDefault();

        const username = document.getElementById('username').value;
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;


        if (username === '' || email === '' || password === '') {
            alert('All fields are required');
            return;
        }
        if (!validateEmail(email)) {
            alert('Invalid email format');
            return;
        }
        if (password.length < 6) {
            alert('Password must be at least 6 characters');
            return;
        }
        /*if (!validateUsername(username)) {
            alert('Username can only contain letters and numbers');
            return;
        }*/

        const data = {
            nickname: username, email: email, password: password
        };

        fetch('/register', {
            method: 'POST', headers: {
                'Content-Type': 'application/json'
            }, body: JSON.stringify(data)
        })
            .then(response => response.json())
            .then(result => {
                if (result.error) {
                    alert(result.error);
                } else {
                    alert(result.message);
                    window.location.href = '/login';
                }
            })
        /*.catch(error => {
            alert(`Error: ${error}`);
            window.location.href = '/';
        });*/
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
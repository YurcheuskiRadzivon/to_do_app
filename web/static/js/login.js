document.addEventListener('DOMContentLoaded', function () {
    const form = document.getElementById('loginForm');

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

        const data = {
            nickname: username, email: email, password: password
        };

        fetch('/login', {
            method: 'POST', headers: {
                'Content-Type': 'application/json'
            }, body: JSON.stringify(data)
        })
            .then(response => response.json())
            .then(result => {
                if (result.error) {
                    alert(result.error);
                } else {
                    const tokenName = "tokenAuth";
                    document.cookie = `${tokenName}=${result.token}; path=/; max-age=86400;`;
                    alert('Login successful!');
                    window.location.href = '/tasks';
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
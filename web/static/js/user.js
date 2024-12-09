document.addEventListener('DOMContentLoaded', function () {
    const menuBtn = document.querySelector('.menu-btn');
    const sidebar = document.querySelector('.sidebar');
    const editBtn = document.querySelector('.edit-btn');
    const saveBtn = document.querySelector('.save-btn');
    const cancelBtn = document.querySelector('.cancel-btn');
    const infoValues = document.querySelectorAll('.info-value');
    const infoInputs = document.querySelectorAll('.info-input');
    const passwordToggle = document.querySelector('.password-toggle');
    const passwordInput = document.querySelector('input[type="password"]');

    // Store original values
    const originalValues = Array.from(infoValues).map(value => value.textContent);

    // Toggle sidebar
    menuBtn.addEventListener('click', function () {
        sidebar.classList.toggle('active');
    });

    // Close sidebar when clicking outside
    document.addEventListener('click', function (e) {
        if (!sidebar.contains(e.target) && !menuBtn.contains(e.target)) {
            sidebar.classList.remove('active');
        }
    });

    // Edit button functionality
    editBtn.addEventListener('click', function () {
        infoValues.forEach(value => value.classList.add('hidden'));
        infoInputs.forEach(input => input.classList.add('active'));
        passwordToggle.classList.add('active');
        saveBtn.classList.add('active');
        cancelBtn.classList.add('active');
        editBtn.style.display = 'none';
    });

    // Save button functionality
    saveBtn.addEventListener('click', async function () {
        try {
            const userData = {
                nickname: infoInputs[0].value,
                email: infoInputs[1].value,
                password: infoInputs[2].value
            };

            const response = await fetch('/user', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(userData)
            })
                .then(response => response.json())
                .then(result => {
                    if (result.error) {
                        alert(result.error);
                    } else {
                        const tokenName = "tokenAuth";
                        document.cookie = `tokenAuth=${result.token}; path=/; max-age=86400;`;

                        window.location.href = '/user';
                    }
                })
        } catch (error) {
            console.error('Error:', error);
            alert('Error updating data');
        }
    });

    // Cancel button functionality
    cancelBtn.addEventListener('click', function () {
        // Restore original values
        infoValues.forEach((value, index) => {
            value.textContent = originalValues[index];
            infoInputs[index].value = index === 2 ? 'password123' : originalValues[index];
        });

        resetEditingState();
    });

    function resetEditingState() {
        infoValues.forEach(value => value.classList.remove('hidden'));
        infoInputs.forEach(input => input.classList.remove('active'));
        passwordToggle.classList.remove('active');
        saveBtn.classList.remove('active');
        cancelBtn.classList.remove('active');
        editBtn.style.display = 'block';
    }

    // Password toggle functionality
    let passwordVisible = false;
    passwordToggle.addEventListener('click', function () {
        passwordVisible = !passwordVisible;
        passwordInput.type = passwordVisible ? 'text' : 'password';
        passwordToggle.textContent = passwordVisible ? '👁️‍🗨️' : '👁️';
    });

    // Logout button functionality
    const logoutBtn = document.querySelector('.logout-btn');
    logoutBtn.addEventListener('click', function () {
        document.cookie = 'tokenAuth=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
        window.location.href = '/';
    });


    const deleteBtn = document.querySelector('.delete-btn');
    deleteBtn.addEventListener('click', async function () {
        if (confirm('Are you sure you want to delete your account? This action cannot be undone.')) {
            try {
                const response = await fetch('/user', {
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });

                if (response.ok) {
                    alert('Account successfully deleted');
                    document.cookie = 'tokenAuth=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
                    window.location.href = '/';
                } else {
                    throw new Error('Error deleting account');
                }
            } catch (error) {
                console.error('Error:', error);
                alert('An error occurred while deleting the account');
            }
        }
    });
});
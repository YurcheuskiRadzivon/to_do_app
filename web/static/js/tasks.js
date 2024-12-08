document.addEventListener('DOMContentLoaded', function () {
    const menuBtn = document.querySelector('.menu-btn');
    const sidebar = document.querySelector('.sidebar');
    const sortBtn = document.querySelector('.sort-btn');
    const sortDropdown = document.querySelector('.sort-dropdown');


    menuBtn.addEventListener('click', function () {
        sidebar.classList.toggle('active');
    });


    document.addEventListener('click', function (e) {
        if (!sidebar.contains(e.target) && !menuBtn.contains(e.target)) {
            sidebar.classList.remove('active');
        }
    });


    sortBtn.addEventListener('click', function (e) {
        e.stopPropagation();
        sortDropdown.classList.toggle('active');
        sortBtn.classList.toggle('active');
    });


    document.addEventListener('click', function (e) {
        if (!sortBtn.contains(e.target)) {
            sortDropdown.classList.remove('active');
            sortBtn.classList.remove('active');
        }
    });


    document.querySelectorAll('.sort-option').forEach(option => {
        option.addEventListener('click', function (e) {
            e.preventDefault();
            const btnText = document.createTextNode('Sort by: ' + this.textContent);
            sortBtn.innerHTML = '';
            sortBtn.appendChild(btnText);

            sortBtn.innerHTML += `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-users-round"><path d="m3 16 4 4 4-4"></path><path d="M7 20V4"></path><path d="m21 8-4-4-4 4"></path><path d="M17 4v16"></path></svg>`;
            sortDropdown.classList.remove('active');


            setTimeout(() => {
                window.location.href = this.href;
            }, 200);
        });
    });
});
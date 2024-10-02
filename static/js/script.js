document.addEventListener('DOMContentLoaded', function() {
    const radioButtons = document.querySelectorAll('input[name="value-radio"]');
    const dynamicLabel = document.getElementById('dynamicLabel');
    const inputField = dynamicLabel.querySelector('input');
    const spanField = dynamicLabel.querySelector('span');

    radioButtons.forEach(radio => {
        radio.addEventListener('change', function() {
            console.log(`Selected value: ${this.value}`); // Добавлено для отладки
            if (this.value === 'value-1') {
                inputField.type = 'text';
                inputField.placeholder = '';
                spanField.textContent = 'Firstname';
            } else if (this.value === 'value-2') {
                inputField.type = 'email';
                inputField.placeholder = '';
                spanField.textContent = 'Email';
            }
        });
    });
});

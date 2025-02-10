// disable perma-button class after its been clicked
document.addEventListener('DOMContentLoaded', function() {
    const actionButtons = document.querySelectorAll('.perma-button');
    actionButtons.forEach(button => {
        button.addEventListener('click', function(e) {
            if (!this.classList.contains('disabled')) {
                this.classList.add('disabled');
                const buttonId = this.getAttribute('data-button-id') || this.textContent.trim();
                localStorage.setItem(`button-${buttonId}`, 'disabled');
                if (!this.hasAttribute('hx-post')) {
                    e.preventDefault();
                }
            }
        });
    });
    actionButtons.forEach(button => {
        const buttonId = button.getAttribute('data-button-id') || button.textContent.trim();
        if (localStorage.getItem(`button-${buttonId}`) === 'disabled') {
            button.classList.add('disabled');
        }
    });
});
// Tab switching
const tabBtns = document.querySelectorAll('.tab-btn');
const tabContents = document.querySelectorAll('.tab-content');

tabBtns.forEach(btn => {
    btn.addEventListener('click', () => {
        // Remove active class from all
        tabBtns.forEach(b => b.classList.remove('active'));
        tabContents.forEach(c => c.classList.remove('active'));

        // Add active class to clicked
        btn.classList.add('active');
        const tabId = btn.dataset.tab;
        document.getElementById(tabId).classList.add('active');
    });
});

// Sync slider with input
const sliders = document.querySelectorAll('.slider');
sliders.forEach(slider => {
    const input = slider.previousElementSibling;

    slider.addEventListener('input', (e) => {
        input.value = e.target.value;
    });

    input.addEventListener('input', (e) => {
        slider.value = e.target.value;
    });
});

// Reset button functionality
const resetBtns = document.querySelectorAll('.reset-btn');
resetBtns.forEach(btn => {
    btn.addEventListener('click', () => {
        const formGroup = btn.closest('.form-group');
        const slider = formGroup.querySelector('.slider');
        const input = formGroup.querySelector('.slider-input');

        if (slider && input) {
            // Reset to default value (you can store default in data attribute)
            const defaultValue = slider.getAttribute('data-default') || slider.min;
            slider.value = defaultValue;
            input.value = defaultValue;
        }
    });
});

// Platform tab switching
const platformTabs = document.querySelectorAll('.platform-tab');
platformTabs.forEach(tab => {
    tab.addEventListener('click', () => {
        const parent = tab.parentElement;
        parent.querySelectorAll('.platform-tab').forEach(t => t.classList.remove('active'));
        tab.classList.add('active');
    });
});

// Form submission
const settingsForm = document.getElementById('settingsForm');
settingsForm.addEventListener('submit', (e) => {
    e.preventDefault();

    // Collect form data
    const formData = new FormData(settingsForm);
    const data = Object.fromEntries(formData);

    console.log('Settings data:', data);

    // Show success message (you can replace with your own notification system)
    alert('Settings saved successfully!');
});
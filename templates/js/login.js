const form = document.getElementById('loginForm');
const messageDiv = document.getElementById('message');
const loginBtn = document.getElementById('loginBtn');

function showMessage(text, type) {
    messageDiv.textContent = text;
    messageDiv.className = `message ${type}`;
    messageDiv.style.display = 'block';

    setTimeout(() => {
        messageDiv.style.display = 'none';
    }, 3000);
}

form.addEventListener('submit', async (e) => {
    e.preventDefault();

    const username = document.getElementById('username').value.trim();
    const password = document.getElementById('password').value;

    if (!username || !password) {
        showMessage('Vui lòng điền đầy đủ thông tin!', 'error');
        return;
    }

    loginBtn.classList.add('loading');
    loginBtn.textContent = '';

    setTimeout(() => {
        loginBtn.classList.remove('loading');
        loginBtn.textContent = 'ĐĂNG NHẬP';

        showMessage('Đăng nhập thành công! Đang chuyển hướng...', 'success');

        setTimeout(() => {
            console.log('Login successful:', { username });
        }, 1500);

    }, 1500);
});

const inputs = document.querySelectorAll('input');
inputs.forEach(input => {
    input.addEventListener('focus', function() {
        this.parentElement.parentElement.style.transform = 'translateX(2px)';
    });

    input.addEventListener('blur', function() {
        this.parentElement.parentElement.style.transform = 'translateX(0)';
    });
});

document.addEventListener('keypress', (e) => {
    if (e.key === 'Enter' && !loginBtn.classList.contains('loading')) {
        form.dispatchEvent(new Event('submit'));
    }
});

// Toggle password visibility
const togglePassword = document.getElementById('togglePassword');
const passwordInput = document.getElementById('password');
const eyeIcon = document.getElementById('eyeIcon');
const eyeOffIcon = document.getElementById('eyeOffIcon');

if (togglePassword) {
    togglePassword.addEventListener('click', function() {
        // Toggle input type
        const type = passwordInput.getAttribute('type') === 'password' ? 'text' : 'password';
        passwordInput.setAttribute('type', type);

        // Toggle icon
        if (type === 'password') {
            eyeIcon.style.display = 'block';
            eyeOffIcon.style.display = 'none';
        } else {
            eyeIcon.style.display = 'none';
            eyeOffIcon.style.display = 'block';
        }
    });
}

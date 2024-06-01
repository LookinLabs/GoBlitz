document.getElementById('signupBtn').addEventListener('click', function() {
    document.getElementById('emailField').classList.remove('hidden');
});

document.getElementById('signinBtn').addEventListener('click', function() {
    document.getElementById('emailField').classList.add('hidden');
});
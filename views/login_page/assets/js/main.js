document.addEventListener('DOMContentLoaded', (event) => {
    document.getElementById('signupBtn').addEventListener('click', function() {
        document.getElementById('signupForm').style.display = 'block';
        document.getElementById('signinForm').style.display = 'none';
    });

    document.getElementById('signinBtn').addEventListener('click', function() {
        document.getElementById('signupForm').style.display = 'none';
        document.getElementById('signinForm').style.display = 'block';
    });
});
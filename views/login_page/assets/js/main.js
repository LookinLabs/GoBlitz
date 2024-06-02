document.addEventListener('DOMContentLoaded', (event) => {
    document.getElementById('signupBtn').addEventListener('click', function() {
        document.getElementById('signupForm').style.display = 'block';
        document.getElementById('signinForm').style.display = 'none';
    });

    document.getElementById('signinBtn').addEventListener('click', function() {
        document.getElementById('signupForm').style.display = 'none';
        document.getElementById('signinForm').style.display = 'block';
    });

    // Handle form submission
    const forms = document.querySelectorAll('#signupForm, #signinForm');
    forms.forEach(form => {
        form.addEventListener('submit', function(event) {
            event.preventDefault(); // Prevent default form submission

            // Submit form using Fetch
            fetch(form.action, {
                method: 'POST',
                body: new FormData(form)
            })
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    // Show banner with error message
                    showBanner(data.error, 'error');
                } else {
                    // Redirect on successful response
                    window.location.href = '/';
                    // Removed the line that shows the success banner
                }
            })
            .catch(error => {
                // Handle network error
                showBanner('Network error', 'error');
            });
        });
    });
});

function showBanner(message, type) {
    const banner = document.createElement('div');
    banner.textContent = message;
    banner.className = type === 'success' ? 'success-banner' : 'error-banner';

    // Create close button
    const closeButton = document.createElement('span');
    closeButton.textContent = 'X';
    closeButton.style.float = 'right';
    closeButton.style.cursor = 'pointer';
    closeButton.addEventListener('click', function() {
        document.body.removeChild(banner);
    });

    // Add close button to banner
    banner.appendChild(closeButton);

    document.body.appendChild(banner);

    // Remove banner after 5 seconds
    setTimeout(function() {
        if (document.body.contains(banner)) {
            document.body.removeChild(banner);
        }
    }, 4000);
}
document.getElementById('generate-form').addEventListener('submit', function(e) {
  e.preventDefault();
  const formData = new FormData(this);
  const formObj = {
    filename: formData.get('filename'),
	width: parseInt(formData.get('width'), 10),
	height: parseInt(formData.get('height'), 10),
	scale: parseFloat(formData.get('scale'))
  };
  fetch(this.action, {
	method: this.method,
	body: JSON.stringify(formObj),
	headers: {
		'Content-Type': 'application/json',
	}
  }).then(response => {
	if (response.ok) {
	  console.log('Request successfully sent to https://localhost:8080/generate');
	  // need to get new.png from the HTTP response
	  response.blob().then(blob => {
	    const url = URL.createObjectURL(blob);
		const a = document.getElementById('download-png-link');
		a.href = url;
		a.download = formObj.filename || 'new.png';
		a.style.display = 'block';
	  });
	  // create a download link that has an href to download new.png
	} else {
	  console.error('Request failed to send to https://localhost:8080/generate');
	}
  }).catch(error => {
	  console.error('Network Error:',error);
  });
});
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
	  return response.blob();
	} else {
	  console.error('Request failed to send to https://localhost:8080/generate');
	}
  }).then(blob => {
	  console.log(blob);
	  const a = document.getElementById('download-png-link');
	  const img = document.getElementById('download-png-preview');
	  if (blob.type === 'application/octet-stream') {
	    const pngBlob = new Blob([blob], { type: 'image/png' });
		const url = URL.createObjectURL(pngBlob);
		img.src = url;
		img.style.display = 'block';
		a.href = url;
		a.download = formObj.filename || 'new.png';
		a.innerHTML = 'Download '+a.download;
		a.style.display = 'block';
	  } else {
		console.error('Received blob from server is not of type "application/octet-stream"');
	  }
  }).catch(error => {
	  console.error('Network Error:',error);
  });
});
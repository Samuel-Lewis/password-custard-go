let last_raw = '';

// ready setups the forms and requests first password
$(document).ready(function () {
	setupForm();

	// GET data from /raw to show
	$('#generate-password').on('click', getRawPassword);
	$('#generate-password').click();

	// Copy on click
	$('#copy-password').on('click', copyToClipboard);

	// Enable tooltips
	$('[data-toggle="tooltip"]').tooltip();
});

// setupForm inis all sliders and form defaults
function setupForm() {
	var w = createSlider('words', 2, 3, 8);
	createSlider('numbers', 1, 1, 8);
	createSlider('symbols', 1, 1, 8);
	createSlider('uppercase', 1, 1, 8);

	$('.slider').each(function (i, obj) {
		obj.noUiSlider.on('update', function () {
			getFormat();
		});
	});

	w.noUiSlider.on('update', function () {
		updateMax();
	});
}

// createSlider makes a noUiSlider and returns it
function createSlider(id, startMin, startMax, max) {
	var slider = document.getElementById(id);
	noUiSlider.create(slider, {
		start: [startMin, startMax],
		connect: true,
		tooltips: [true, true],
		format: wNumb({
			decimals: 0,
			thousand: '',
			suffix: '',
		}),
		step: 1,
		animate: true,
		range: {
			'min': 0,
			'max': max
		},
	});

	return slider;
}

// updateMax scales all length dependent sliders (.capped) to be capped at total number of words
function updateMax() {
	var w = document.getElementById('words').noUiSlider.get();
	var m = parseInt(w[1]);
	m = Math.max(m, 1);
	$('.capped').each(function (i, obj) {
		obj.noUiSlider.updateOptions({
			range: {
				'min': 0,
				'max': m,
			}
		})
	});
}

// getFormat reads form elements and generates formatting code (and shows it)
function getFormat() {
	var cust = document.getElementById('custom');
	var f = '';

	if (cust.value != '') {
		f = cust.value;
	} else {
		$('.slider').each(function (i, obj) {
			var vals = obj.noUiSlider.get();
			if (vals[0] != '0' || vals[1] != '0') {
				var s = obj.id;

				// Optional quantity args formatting
				if (vals[0] == vals[1]) {
					if (vals[0] != '1') {
						s += ':' + vals[0];
					}
				} else {
					s += ':' + vals[0] + ':' + vals[1];
				}

				f += s + ',';
			}
		});
		f = f.slice(0, -1);
	}

	document.getElementById('format').innerHTML = f;
	return f;
}

// getRawPassowrd sends ajax request to /raw, with formatting options if present
function getRawPassword() {
	$('#copy-password').attr('data-original-title', 'Copy');

	var f = getFormat();
	if (f != '') {
		f = '?q=' + f;
	}

	$.ajax({
		url: '/raw' + f,
		method: 'GET',
		success: function (data) {
			pass = $('<div/>').text(data).html();
			// Send to heading
			$('#password-field').html(pass);

			last_raw = data;
			logPass(pass);
		},
	});
}

// copyToClipboard copies current password to client clipboard
function copyToClipboard() {
	// Update tooltip
	$('#copy-password').attr('data-original-title', 'Copied!')
		.tooltip('show');

	// Create temp textarea, put text in it, copy from that
	var textArea = document.createElement('textarea');
	textArea.value = last_raw;
	document.body.appendChild(textArea);

	textArea.select();
	try {
		document.execCommand('copy');
	} catch (err) {
		alert('Browser does not support copy and paste automation :(');
	}
	document.body.removeChild(textArea);
}

// loadPass records and shows last 10 passwords to history box
const history = [];

function logPass(pass) {
	$('#history').html(history.join('</br>'));

	history.unshift(pass);

	if (history.length > 10) {
		history.pop();
	}
}
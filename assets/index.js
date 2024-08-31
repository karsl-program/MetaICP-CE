function showPopup() {
    document.getElementById('popup').style.display = 'block';
    document.getElementById('popup').style.animation = 'change 0.2s';
    document.body.style.background = 'rgb(0, 0, 0, 0.3)'
    var x = document.getElementsByClassName("text")
    for (var i = 0; i < x.length; i++) {
        x[i].style.color = '#999'
    }
}

function hidePopup() {
    document.body.style.background = ''
    var x = document.getElementsByClassName("text")
    for (var i = 0; i < x.length; i++) {
        x[i].style.color = ''
    }
    document.getElementById('popup').style.display = 'none';
}
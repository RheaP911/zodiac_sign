// Star
window.addEventListener('mousemove', function(e) {
    var arr = [1, 0.9, 0.8, 0.5, 0.2];

    arr.forEach(function(i) {
        var x = (1 - i) * 75;
        var star = document.createElement('div');

        star.className = 'star';
        star.style.top = e.pageY + Math.round(Math.random() * x - x / 2) + 'px';
        star.style.left = e.pageX + Math.round(Math.random() * x - x / 2) + 'px';

        document.body.appendChild(star);

        window.setTimeout(function() {
            document.body.removeChild(star);
        }, Math.round(Math.random() * i * 600));
    });
}, false); 
    

// Sidebar
const body = document.querySelector('body'),
    sidebar = body.querySelector('nav'),
    searchBtn = body.querySelector(".search-box"),
    modeSwitch = body.querySelector(".toggle-switch"),
    modeText = body.querySelector(".mode-text");

const logoImg = document.querySelector('.image img');
let wrapper = document.querySelector('.dashboard');

sidebar.addEventListener("mouseenter", () => {
    logoImg.src = "/static/resources/logo_img.png";
    wrapper.style.left = "250px";
    wrapper.style.width = "calc(100% - 250px)";
});

sidebar.addEventListener("mouseleave", () => {
    logoImg.src = "/static/resources/logo-gwhite.png";
    wrapper.style.left = "88px";
    wrapper.style.width = "calc(100% - 88px)";
    
    wrapper.addEventListener("mouseenter", () => {
        wrapper.style.left = "88px";
        wrapper.style.width = "calc(100% - 88px)";
    });
});

// Search

// Get the input element
var input = document.getElementById("searchInput");
// Get the menu links container
var menuLinks = document.getElementById("zodiac-buttons");
// Get all the menu links
var links = menuLinks.getElementsByTagName("li");

// Add event listener for input event
input.addEventListener("input", function() {
    // Get the search query
    var searchQuery = input.value.toLowerCase();
    // Loop through all menu links
    for (var i = 0; i < links.length; i++) {
        var linkText = links[i].textContent.toLowerCase();
        // Check if the link text contains the search query
        if (linkText.includes(searchQuery)) {
            // Show the menu link
            links[i].style.display = "block";
        } else {
            // Hide the menu link
            links[i].style.display = "none";
        }
    }
});


//Active sidebar buttons
// var btnContainer = document.getElementById("zodiac-buttons");

// var btns = btnContainer.getElementsByClassName("zodiac-btn");

// for (var i = 0; i < btns.length; i++) {
//     btns[i].addEventListener("click", function() {
//         var current = document.getElementsByClassName("active");
//         current[0].className = current[0].className.replace(" active", "");

//         this.className += " active";
//     });
// }



"use strict";

/* ====== Define JS Constants ====== */
const sidebarToggler = document.getElementById('docs-sidebar-toggler');
const sidebar = document.getElementById('docs-sidebar');
const sidebarLinks = document.querySelectorAll('#docs-sidebar .scrollto');

let navScroll = false;
let currentNavItemGroup = null;

/* ===== Responsive Sidebar ====== */

window.onload = function () {
	responsiveSidebar();
};

window.onresize = function () {
	responsiveSidebar();
};


function responsiveSidebar() {
	let w = window.innerWidth;
	if (w >= 1200) {
		// if larger 
		sidebar.classList.remove('sidebar-hidden');
		sidebar.classList.add('sidebar-visible');

	} else {
		// if smaller
		sidebar.classList.remove('sidebar-visible');
		sidebar.classList.add('sidebar-hidden');
	}
};

sidebarToggler.addEventListener('click', () => {
	if (sidebar.classList.contains('sidebar-visible')) {
		sidebar.classList.remove('sidebar-visible');
		sidebar.classList.add('sidebar-hidden');

	} else {
		sidebar.classList.remove('sidebar-hidden');
		sidebar.classList.add('sidebar-visible');
	}
});

/* ===== Smooth scrolling ====== */
/*  Note: You need to include smoothscroll.min.js (smooth scroll behavior polyfill) on the page to cover some browsers */
/* Ref: https://github.com/iamdustan/smoothscroll */


let currentHash = window.location.hash;
let currentSection;

let currentNavItem = $('a[href*="' + window.location.hash + '"]');
let currentNavSection = $('a[href*="' + window.location.hash + '"]').parents("li").prevAll(".section-title:first");

if (currentNavSection.length === 0) {
	currentNavSection = $(".section-title:first");
}

let currentNavSectionHash = currentNavSection.find("a");

if (currentNavSectionHash.length !== 0) {
	currentNavSectionHash = currentNavSectionHash[0].hash
}

if (currentHash === "") {
	currentSection = $(".section-title:first");
	currentNavSectionHash = currentSection.find("a")[0].hash;
}

if (currentNavSectionHash === "#golang-application") {
	currentNavSection.find("span img").attr("src", "assets/images/golang-logo-white.svg");
}

currentNavItem.first().addClass("active");
currentNavSection.addClass("active");

sidebarLinks.forEach((sidebarLink) => {
	sidebarLink.addEventListener('click', (e) => {
		navScroll = true;
		e.preventDefault();

		var target = sidebarLink.getAttribute("href").replace('#', '');

		currentNavItemGroup = $(e.target).parent("li").prevAll(".section-title");

		$(".section-title").removeClass("active");
		$(".nav-link").removeClass("active");
		$(".nav-item").removeClass("active");

		if ($(e.target).parents("li").hasClass("section-title")) {
			$(e.target).parents("li").addClass("active")
		} else {
			$(e.target).addClass("active");
			currentNavItemGroup.first().addClass("active");
		}

		window.location.hash = sidebarLink.getAttribute("href");

		if ($(".section-items").find("a[href*='#golang-application']").parent("li").hasClass("active")) {
			$(".section-items").find("a[href*='#golang-application']").find("span img").attr("src", "assets/images/golang-logo-white.svg");
		} else {
			$(".section-items").find("a[href*='#golang-application']").find("span img").attr("src", "assets/images/golang-logo-green.svg");
		}

		document.getElementById(target).scrollIntoView({ behavior: 'smooth' });

		//Collapse sidebar after clicking
		if (sidebar.classList.contains('sidebar-visible') && window.innerWidth < 1200) {
			sidebar.classList.remove('sidebar-visible');
			sidebar.classList.add('sidebar-hidden');
		}
	});
});

const docLinks = document.querySelectorAll('.docs-content a');
docLinks.forEach((a) => {
	if (a.host === window.location.host) {
		a.addEventListener('click', (e) => {
			e.preventDefault();

			$(".section-title").removeClass("active");
			$(".nav-link").removeClass("active");
			$(".nav-item").removeClass("active");

			let navLink = $('#docs-sidebar a[href="' + a.hash + '"]');
			let currentNavItemGroup = navLink.parent("li").prevAll(".section-title");

			if (navLink.parent().hasClass("section-title")) {
				navLink.parent().addClass("active")
			} else {
				navLink.parent().addClass("active");
				currentNavItemGroup.first().addClass("active");
			}

			if ($(".section-items").find("a[href*='#golang-application']").parent("li").hasClass("active")) {
				$(".section-items").find("a[href*='#golang-application']").find("span img").attr("src", "assets/images/golang-logo-white.svg");
			} else {
				$(".section-items").find("a[href*='#golang-application']").find("span img").attr("src", "assets/images/golang-logo-green.svg");
			}

			document.getElementById(a.hash.replace('#', '')).scrollIntoView({ behavior: 'smooth' });

			//Collapse sidebar after clicking
			if (sidebar.classList.contains('sidebar-visible') && window.innerWidth < 1200) {
				sidebar.classList.remove('sidebar-visible');
				sidebar.classList.add('sidebar-hidden');
			}
		});
	}
});

/* ====== SimpleLightbox Plugin ======= */
/*  Ref: https://github.com/andreknieriem/simplelightbox */

var lightbox = new SimpleLightbox('.simplelightbox-gallery a', {/* options */ });


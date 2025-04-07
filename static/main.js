(function ($) {
  "use strict";

  // Toggle sidebar menu
  $(".menu-toggle").on("click", function (e) {
    e.preventDefault();
    $("#sidebar-wrapper").toggleClass("active");
    $(this).toggleClass("active");
    $(".menu-toggle > .fa-bars, .menu-toggle > .fa-times").toggleClass("fa-bars fa-times");
  });

  // Smooth scrolling for anchor links
  $('a.js-scroll-trigger[href*="#"]:not([href="#"])').on("click", function () {
    const { pathname, hostname, hash } = this;
    if (
      location.pathname.replace(/^\//, "") === pathname.replace(/^\//, "") &&
      location.hostname === hostname
    ) {
      let target = $(hash);
      target = target.length ? target : $("[name=" + hash.slice(1) + "]");
      if (target.length) {
        $("html, body").animate(
          { scrollTop: target.offset().top },
          1000,
          "easeInOutExpo"
        );
        return false;
      }
    }
  });

  // Close sidebar when a scroll trigger link is clicked
  $("#sidebar-wrapper .js-scroll-trigger").on("click", function () {
    $("#sidebar-wrapper").removeClass("active");
    $(".menu-toggle").removeClass("active");
    $(".menu-toggle > .fa-bars, .menu-toggle > .fa-times").toggleClass("fa-bars fa-times");
  });

  // Show/hide scroll-to-top button on scroll
  $(document).on("scroll", function () {
    $(".scroll-to-top").fadeToggle($(this).scrollTop() > 100);
  });
})(jQuery);

$(document).ready(function () {
  var pathArr = window.location.pathname.split("/").slice(0, 2);
  pathArr = pathArr.filter(p => p);
  pathArr = pathArr.filter(p => !p.includes(".html"));
  var homePath = "/"+pathArr.join("");
  $(".homelink").attr("href", homePath)
  
  $(".code-copy-button").each(function (index) {
    const copyButton = $(this)[0];
    let codeBoxPos = {}
    codeBoxPos.top = $(this).siblings().first().position().top;
    codeBoxPos.right = $(this).siblings().first().position().left + $(this).siblings().first().width();
    codeBoxPos.bottom = $(this).siblings().first().position().top + $(this).siblings().first().height();
    codeBoxPos.left = $(this).siblings().first().position().left;
    if ($(this).parents("pre").length > 0) {
      codeBoxPos.right = $(this).parents("pre").first().position().left + $(this).parents("pre").first().width() - 10;
    }
    $(this).siblings().first().css("padding-right", "30px")
    $(this).css("position", "absolute");
    $(this).css("top", codeBoxPos.top + 2);
    $(this).css("left", codeBoxPos.right + 11);
    copyButton.addEventListener('click', (event) => {
      const content = $(this).siblings().first().text();
      const self = $(this);
      navigator.clipboard.writeText(content).then(function () {
        // Promise Success
        self.removeClass("fa-copy");
        self.addClass("fa-check");

        setTimeout(() => {
          self.removeClass("fa-check");
          self.addClass("fa-copy");
        }, 2000)
      }, function () {
        // Promise Error
      });
    });
  });

  $(".project-section").each(function (index) {
    const card = $(this);
    const link = $(this).find("a")[0];
    card.css("cursor", "pointer");

    card[0].addEventListener('click', (event) => {
      link.click()
    });

    const url = new URL(link.href);

    if (window.location.pathname.slice(1) === url.pathname.slice(1)) {
      $(this).css("background", "#eff6ef");
    }

  });
});
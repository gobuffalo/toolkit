require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

const gh = "https://github.com";
const rawGh = "https://raw.githubusercontent.com";

$(() => {
  $("#repo-deets #readme-tab").tab("show");

  $("img[src^='https://github.com']").each((i, el) => {
    var img = $(el);
    var src = img.attr("src");
    src = src.replace(gh, rawGh);
    src = src.replace("blob/", "");
    img.attr("src", src);
    // TODO: get .svg displaying
    if (src.endsWith(".svg")) {
      img.remove();
    }
  });

});

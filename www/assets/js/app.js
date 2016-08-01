var mainView = null;
var sideView = null;

$(document).ready(function() {
	mainView = $("div#main-view");
	sideView = $("div#side-view");
	init();
});

function init() {
	$('a#tags').click(function() {
		$.ajax({
			url: "/api/tags",
			method: "get",
			success: function(r) {
				createListTags(r);
			}
		});
	});

	$(document).on('click', 'a.tag', function(e) {
		key = e.target.id;
		alert(key);
	});
}

function createListTags(tags) {
	mainView.empty()
	$.each(tags, function(k, v) {
		a = $('<a href="#" class="btn btn-primary btn-xs tag">');
		a.attr('id', v._key);
		a.append(v.name);
		mainView.append(a);
		mainView.append('&nbsp;')
	});
}

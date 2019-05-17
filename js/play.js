function initPlayer(playerId){
  'use strict';
  var langMap = {
	rb: 'ruby', go: 'go', py: 'python', java: 'java', md: 'markdown', js: 'js'
  };

  var el = document.getElementById(playerId);

  el.innerHTML = `
<ul class="nav nav-tabs md-tabs" id="myPlayerTabs" role="tablist"></ul>
<div class="player-code-body tab-content card" id="myPlayerCode"></div>

<div class="player-code-result ml-3">
<p>Result</p>
<pre><code class="remark-code"></code></pre>
</div>

<div class="player-control mt-3">
<button class="btn btn-raised" onclick="player.run();">Run</button>
<button class="btn btn-raised" onclick="player.close();">Close</button>
</div>`;

  function getSrc(s) {
	return fetch(s, {cache: 'no-cache'}).then(function(res) {
	  return res.text();
	});
  };

  function selectTab(e) {
    console.log(e);
	$(e).tab('show');
  };

  function resultText(txt) {
	  document.querySelector('.player-code-result > pre > code').innerHTML = txt;
  };

  function updateTab(tabs, codes) {
	document.getElementById('myPlayerTabs').innerHTML = tabs;
	document.getElementById('myPlayerCode').innerHTML = codes;
	resultText('');
  };
  
  function open() {
	var tabs = '';
	var calls = [];
	Array.from(arguments, function(f, i) {
	  var elId = f.replace(/[\/\.]/g, '_');
	  var nav = `<li class="nav-item">
    <a class="nav-link" id="${elId}-tab" data-toggle="tab" href="#${elId}" role="tab" aria-controls="${elId}"
      aria-selected="true" onclick="player.tab(this.parentNode);">${f}</a>
  </li>`;
	  tabs += nav;
	  calls.push(getSrc(f).then(function(src) {
		var contentId = f.replace(/[\/\.]/g, '_');
		var lang = langMap[f.split('.').pop()];
		return `<div class="tab-pane fade show" id="${contentId}" role="tabpanel" aria-labelledby="${contentId}-tab" code-lang="${lang}">
<pre class="mt-2 ml-3 mb-2"><code class="remark-code">${src}</code><textarea style="display:none;">${src}</textarea></pre>
</div>`;
	  }));
	});
	  
	Promise.all(calls).then(function(values) {
	  var codes = values.join('');
	  updateTab(tabs, codes);
	  $('#player').show();
	}).then(function() {
      selectTab(document.querySelector('#player .nav-item'));
    });
  };

  function dispose() {
	updateTab('', '');
	$('#player').hide();
  };

  
  function run() {
	var tab = document.querySelector('.tab-pane.active');
	var lang = tab.getAttribute('code-lang');
	var code = tab.querySelector('pre > textarea').value;
	resultText('');

	fetch('/run', {
	  method:'POST',
	  cache: 'no-cache',
	  headers: {"Content-Type": "application/json"},
	  body: JSON.stringify({lang:lang,code:code})
	}).then(function(res) {
	  return res.text(); 
	}).then(function(result) {
	  console.log(result);
	  resultText(result);
	});
  };

  var player = {
	tab: selectTab,
	open: open,
	close: dispose,
	run: run
  };
  window.player = player;
  dispose();
}

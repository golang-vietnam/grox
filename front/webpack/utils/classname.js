var regSource = /exports\.push\((.*)\);/;
var regLocals = /exports\.locals = ((.|\n)+);/;
var charMaps = [];
var notFirst = initNotFirst();

function initNotFirst() {
	var res = {};
	var notFirstChars = '-0123456789';
	for (var i = 0; i < notFirstChars.length; i++) {
		var ch = notFirstChars[i];
		res[ch] = true;
	}

	return res;
}

function initCharmap() {
	var chars = '_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-'.split('');

	var charMap = {};
	var i, ch, cur = '@', next;
	while (chars.length) {
		i = Math.floor(Math.random()*chars.length);
		next = chars[i];
		chars.splice(i, 1);

		charMap[cur] = next;
		cur = next;
	}
	charMap[next] = undefined;

	return charMap;
}

function inc(ch, position) {
	var charMap = charMaps[position];
	if (!charMap) {
		charMap = initCharmap();
		charMaps[position] = charMap;
	}

	if (!ch) ch = '@';
	if (position) return charMap[ch];

	do {
		ch = charMap[ch];
	} while (notFirst[ch]);
	return ch;
}

function nextId(id) {
	if (!id) return inc();

	var ids = id.split('');
	for (var i=0; true; i++) {
		var ch = inc(ids[i], i);
		if (ch) {
			ids[i] = ch;
			break;
		}

		ids[i] = inc(undefined, i);
	}

	return ids.join('');
}

var classMap = [];
var __curId;

module.exports = function rep(content) {
	try {
	var sourceMatch = content.match(regSource);
	var localsMatch = content.match(regLocals);
	if (!sourceMatch || !localsMatch) {
		return content;
	}

	var source = sourceMatch[1];
	var locals = JSON.parse(localsMatch[1]);

	for (var key in locals) {
		var className = locals[key];
		if (!className) continue;

		var newName = nextId(__curId);
		classMap[className] = newName;
		__curId = newName;
	}

	var newContent = content.replace(/[-_A-z0-9]{18,28}/g, function(oldName) {
		return classMap[oldName] || oldName;
	});

	return newContent;

	} catch(e) {
		console.error('ERROR classname-loader', e);
		return content;
	}
}

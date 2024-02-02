(function() {
  window.__onThemeChange = function() {};
  function setTheme(newTheme) {
    window.__theme = newTheme;
    preferredTheme = newTheme;
    document.body.className = newTheme;
    window.__onThemeChange(newTheme);
  }
  var preferredTheme;
  try {
    preferredTheme = localStorage.getItem('theme');
  } catch (err) { }
  window.__setPreferredTheme = function(newTheme) {
    setTheme(newTheme);
    try {
      localStorage.setItem('theme', newTheme);
    } catch (err) {}
  }
  setTheme(preferredTheme || 'dark');
})();
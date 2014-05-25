<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1">
<title id="title_content">ExtTop - Desktop Sample App</title>
    <link rel="stylesheet" type="text/css" href="/ExtJs/resources/css/ext-all.css" />
    <link rel="stylesheet" type="text/css" href="/ExtJs/desktop/css/desktop.css" />
    <!-- GC -->
    <!-- LIBS -->
    <script type="text/javascript" src="/ExtJs/adapter/ext/ext-base.js"></script>
    <!-- ENDLIBS -->
    <script type="text/javascript" src="/ExtJs/ext-all-debug.js"></script>
    <!-- DESKTOP -->
    <script type="text/javascript" src="/ExtJs/desktop/js/StartMenu.js"></script>
    <script type="text/javascript" src="/ExtJs/desktop/js/TaskBar.js"></script>
    <script type="text/javascript" src="/ExtJs/desktop/js/Desktop.js"></script>
    <script type="text/javascript" src="/ExtJs/desktop/js/App.js"></script>
    <script type="text/javascript" src="/ExtJs/desktop/js/Module.js"></script>
    <script type="text/javascript" src="/ExtJs/desktop/sample.js"></script>
</head>
<body scroll="no">
<div id="x-desktop">
    <a href="http://extjs.com" target="_blank" style="margin:5px; float:right;"><img src="/ExtJs/desktop/images/powered.gif" /></a>

    <dl id="x-shortcuts">
        <dt id="grid-win-shortcut">
            <a href="#"><img src="/ExtJs/desktop/images/s.gif" />
            <div>Grid Window</div></a>
        </dt>
        <dt id="acc-win-shortcut">
            <a href="#"><img src="/ExtJs/desktop/images/s.gif" />
            <div>Accordion Window</div></a>
        </dt>
        <dt id="accs-win-shortcut">
            <a href="#"><img src="/ExtJs/desktop/images/s.gif" />
            <div>这是我测试的</div></a>
        </dt>
    </dl>
</div>
<div id="ux-taskbar">
    <div id="ux-taskbar-start"></div>
    <div id="ux-taskbuttons-panel"></div>
    <div class="x-clear"></div>
</div>
</body>
</html>
<script type="text/javascript">
    document.getElementById('accs-win-shortcut').onclick = function(){
        document.getElementById('title_content').innerHTML = "你有新的消息，，，，";
    }
</script>

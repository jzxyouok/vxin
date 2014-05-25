<!DOCTYPE html>
<html>
<head>
    <title>this is a demo</title>
    <link rel="stylesheet" type="text/css" href="/ExtJs/resources/css/ext-all.css">
    <script src="/ExtJs/adapter/ext/ext-base-debug.js"></script>
    <script src="/ExtJs/ext-all-debug.js"></script>
    <script src="/ExtJs/src/locale/ext-lang-zh_CN.js"></script>
</head>
<body>
<div id="panel" style="float:left"></div>
<div id="table" style="float:left"></div>
</body>
</html>
<script type="text/javascript">
    Ext.onReady(function(){
        // new Ext.Window({title:'hello',width:400,height:100,html:'<h1>你好，这是我测试的</h1>'}).show();
        // Ext.MessageBox.confirm("helo","你好啊");
        // var panel = new Ext.Panel({title:'你好，这是网格',width:200,height:200,html:'this is a gridpanel',renderTo:'panel'});
        // var table = new Ext.TabPanel({
        //         title:'这是一个表格',width:200,height:200,
        //         items:[
        //             {title:'面板1',heigth:30,html:'面板1'},
        //             {title:'面板2',height:30,html:'面板2'}
        //             ],
        //         renderTo:'table'});
       


/*******************弹出框***************/
    // Ext.Msg.show({
    //     title : 'alert_one',
    //     msg : 'say hello ',
    //     buttons :{
    //         yes:"请你确认一下？",
    //         no :true,
    //         cancel :true
    //     },
    //     icon:'milton-icon',
    //     fn:function(eventype){
    //         switch(eventype){
    //             case 'yes' :
    //                 Ext.Msg.prompt('this is a editbox', 'where is it ?', function(eventtype, content){
    //                     alert(eventtype+'--'+content);
    //                 },function(){},true,'吴赐有');

    //             break ;
    //         }
    //     }
    // })

/*************表单***************************************/
    var movie_form = new Ext.FormPanel({
        url:'mvoe-submit.php',
        renderTo: document.body,
        frame:true,
        title:'this is a editform',
        width:300,
        items:[
            {xtype:'textfield',vtype:'email',fieldLabel:'Title',name:'title',allowBlank:false},
            {xtype:'textfield',fieldLabel:'Director',name:'Director'},
            {xtype:'datefield',fieldLabel:'Released',name:'Released',width:140,disabledDays:[1,2,3]}

        ]

    });

    new Ext.FormPanel({
        url :'fewfew.php',
        renderTo : document.body,
        frame : true, 
        title :'this is a form ',
        width:400,
        items:[
            {xtype:'textfield',fieldLabel:'name'}
        ]
    });

//         Ext.create('Ext.Button', {
//     text: 'Click me',
//     renderTo: Ext.getBody(),
//     handler: function() {
//         alert('You clicked the button!');
//     }
// });
    })
</script>
var code;
function change(){
    var arrays=new Array(
    '1','2','3','4','5','6','7','8','9','0',
    'a','b','c','d','e','f','g','h','i','j',
    'k','l','m','n','o','p','q','r','s','t',
    'u','v','w','x','y','z',
    'A','B','C','D','E','F','G','H','I','J',
    'K','L','M','N','O','P','Q','R','S','T',
    'U','V','W','X','Y','Z');
     code='';
     for(var i=0;i<4;i++){
         var r=parseInt(Math.random()*arrays.length);
         code+=arrays[r];
     }
     document.getElementById("code").innerHTML=code;
 }

function check(){
    var icode=document.getElementById("vcode").value;
    if(icode==''){
        document.getElementById("cw2").className="wrong";
        document.getElementById("cw2").innerText="验证码不能为空!";
        return false;
    }
    else{
        if(icode!=code){
            document.getElementById("cw2").className="wrong";
            document.getElementById("cw2").innerText="请输入正确的验证码";
            $("#vcode").val("");
            return false;
        }else{
            document.getElementById("cw2").innerText=" ";
            return true;
        }
    }
}  
 
function checkname() {
    var name = document.getElementById("user").value;
    if(name==''){ 
        document.getElementById("cw").className="wrong";
        document.getElementById("cw").innerText="用户名不能为空!";
        return false;
    } else {
        document.getElementById("cw").innerText=" ";
        return true 
    }
}
function checkpass() {
    var pass = document.getElementById("password").value;
    var reg = /^[0-9a-zA-Z]{6,12}$/;
    if(pass==''){ 
        document.getElementById("cw1").className="wrong";
        document.getElementById("cw1").innerText="密码不能为空!";
        
        return false;
    }
     else{
        if (reg.test(pass)) {
            document.getElementById("cw1").innerText=" ";
            return true;
        } else {
            document.getElementById("cw1").className="wrong";
            document.getElementById("cw1").innerText="密码长度为6—12位";
            $("#password").val("");
            return false;
           }
        }
}

function checkall(){
    var now=new Date();
    var hour=now.getHours();
    var minute=now.getMinutes();
    var second=now.getSeconds();
    var week=now.getDay();
    var a=['日','一','二','三','四','五','六'];
    var str="现在是："+hour+":"+minute+":"+second+",星期"+a[week];
    if(checkname()==true&&checkpass()==true&&check()==true){
        alert(str);
        return true
    }else{
        alert("登录失败！");
        location.href="login.html";
        return false;
    }
}
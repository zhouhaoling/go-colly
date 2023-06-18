function checkname() {
    var name = document.getElementById("username").value;
    if(name==''){ 
        document.getElementById("cw").className="wrong";
        document.getElementById("cw").innerText="请输入用户名";
        return false
    } else{
        document.getElementById("cw").innerText="";
        return true
    }
}
function checkpass() {
    var pass1 = document.getElementById("pass1").value;
    var reg = /^[0-9a-zA-Z_-]{6,12}$/;
 
    if (reg.test(pass1)) {
        document.getElementById("cw1").innerText=" ";
        return true;
    } else {
        document.getElementById("cw1").className="wrong";
        document.getElementById("cw1").innerText="密码长度为6—12位";
        return false;
    }
 
}

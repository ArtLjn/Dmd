"use strict";(self["webpackChunkvue_template"]=self["webpackChunkvue_template"]||[]).push([[570],{1880:function(e,a,o){o.r(a),o.d(a,{default:function(){return w}});var n=o(6768),s=o(5130),t=o.p+"static/img/login.a678b768.jpg";const i=e=>((0,n.Qi)("data-v-643ec988"),e=e(),(0,n.jt)(),e),r={id:"login"},l={class:"container",id:"container"},c={class:"form-container sign-in-container"},d={action:"#"},u=i((()=>(0,n.Lk)("h1",null,"个人健康系统",-1))),p=i((()=>(0,n.Lk)("div",{class:"overlay-container"},[(0,n.Lk)("div",{class:"overlay"},[(0,n.Lk)("div",{class:"overlay-panel overlay-right"},[(0,n.Lk)("img",{class:"logo",src:t,alt:""})])])],-1)));function m(e,a,o,t,i,m){return(0,n.uX)(),(0,n.CE)("div",r,[(0,n.bF)(s.eB,{appear:"",name:"opacitytrans"},{default:(0,n.k6)((()=>[(0,n.Lk)("div",l,[(0,n.Lk)("div",c,[(0,n.Lk)("form",d,[u,(0,n.Lk)("div",null,[(0,n.bo)((0,n.Lk)("input",{type:"text",placeholder:"用户名","onUpdate:modelValue":a[0]||(a[0]=e=>i.loginForm.username=e)},null,512),[[s.Jo,i.loginForm.username]]),(0,n.bo)((0,n.Lk)("input",{type:"password",placeholder:"密码","onUpdate:modelValue":a[1]||(a[1]=e=>i.loginForm.password=e)},null,512),[[s.Jo,i.loginForm.password]]),(0,n.Lk)("div",{class:"button",onClick:a[2]||(a[2]=e=>m.handleLogin())},"登录")])])]),p])])),_:1})])}o(4114);var g=o(9325),h=o(5720);const v=e=>h.A.post("/api/user/login",e),k=()=>h.A.get("/api/user/verifyToken");var L={data(){return{loginForm:{username:"",password:""}}},methods:{handleLogin(){""!==this.loginForm.username&&""!==this.loginForm.password?v(this.loginForm).then((e=>{const a=e.data.data[0];localStorage.setItem("token",a),this.$message.success("登录成功"),g.A.push("/home")})):this.$message.warning("请完整填写表单")}},mounted(){k().then((()=>{g.A.push("/home")}))}},f=o(1241);const F=(0,f.A)(L,[["render",m],["__scopeId","data-v-643ec988"]]);var w=F}}]);
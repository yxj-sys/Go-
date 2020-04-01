package utils

import (
	"fmt"
)
//收支系统结构体


type FamilyAccount struct {
	admin		string//软件账号
	password	string//软件密码
	key 		string//保存接受用户输入
	loop 		bool//控制for循环用户界面
	balance 	float64//账户余额
	money		float64//每次收支金额
	note		string//每次收支说明
	flag 		bool//记录是否有收支行为
	details		string//收支字段
	account 	string//转账账户	
}

//编写要给工厂模式的构造方法，返回一个*FamilyAccount的实例
func NewFamilyAccount() *FamilyAccount {
	return	&FamilyAccount{
	admin : "admin",
	password : "123456",
	key  : "",
	loop : true,
	balance : 10000.0,
	money : 0.0,
	note : "",
	flag :  false,
	details : "类型\t收支\t账户余额\t说	明",
	account : "",
	}
}
//登录软件方法
func (this *FamilyAccount) Signin() {
	fmt.Println("请输入用户名:")
	fmt.Scanln(&this.admin)
	fmt.Println("请输入密码:")
	fmt.Scanln(&this.password)
	for {
		if this.admin =="admin" && this.password == "123456" {
			fmt.Println("登录成功！")
		}else {
			fmt.Println("账号或密码错误，请重新输入!")
			this.Signin()
		}
		break
	}
	
}


//收支明细记录方法
func (this *FamilyAccount) showDetails() {
		fmt.Println("----------------当前收支记录---------------------")
			if	this.flag {
				fmt.Println(this.details)
			}else {
				fmt.Println("当前没有收支记录，先记录一笔吧！")
			}
}

//登记收入方法
func (this *FamilyAccount) incoun() {
	fmt.Println("本次收入金额：")
		for {
			fmt.Scanln(&this.money)
			if this.money <= 0 {
				fmt.Println("输入错误，请重新输入")
			}else {
				this.balance += this.money //修改账户余额
				break
			}
		}
		fmt.Println("本次收入说明：")
		fmt.Scanln(&this.note)
		//将本次收入情况拼接到details变量
		this.details += fmt.Sprintf("\n收入\t%v\t%v\t	%v",this.money,this.balance,this.note)
		this.flag = true
}


//登记支出方法
func (this *FamilyAccount) expend() {
	fmt.Println("本次支出金额：")
		for {
			fmt.Scanln(&this.money)
			if this.balance - this.money < 0 {
				fmt.Println("余额不足，请重新输入")
			}else {
				this.balance -= this.money //修改账户余额
				break
			}
		}
		fmt.Println("本次支出说明：")
		fmt.Scanln(&this.note)
		//将本次收入情况拼接到details变量
		this.details += fmt.Sprintf("\n支出\t	%v\t%v\t	%v",this.money,this.balance,this.note)
		this.flag = true
}

//转账方法
func (this *FamilyAccount) transferaccount() {
	fmt.Println("请输入您要转账的账号:")
	fmt.Scanln(&this.account)
	fmt.Println("请输入您要转账的金额:")
	for {
			fmt.Scanln(&this.money)
			if this.balance - this.money < 0 {
				fmt.Println("余额不足，请重新输入")
			}else {
			this.balance -= this.money //修改账户余额
			break
			}
		}
		fmt.Println("本次转账说明：")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n转账\t%v\t%v\t	%v %v",this.money,this.balance,this.account,this.note)
		this.flag = true
}

//退出系统方法
func (this *FamilyAccount) exit() {
	var i string //接收退出系统变量
	fmt.Println("是否退出系统？Y/N")
	fmt.Scanln(&i)
	for{
		if i == "Y" {
			this.loop = false
			break
		}else if i == "N" {
			this.loop = true
			break
		}else {
			fmt.Println("输入错误！")
			break
		}
	}
}

//主菜单方法
func (this *FamilyAccount) MainMenu() {
	for {
		this.Signin()
		fmt.Println("---------------家庭收支记账系统-------------------")
		fmt.Println("---------------1.收支明细")
		fmt.Println("---------------2.登记收入")
		fmt.Println("---------------3.登记支出")
		fmt.Println("---------------4.转账")
		fmt.Println("---------------5.退出系统")
		fmt.Println("请选择（1-5）")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1": 
				this.showDetails()
			case "2": 
				this.incoun()
			case "3": 
				this.expend()
			case "4": 
				this.transferaccount()
			case "5":
				this.exit()
		}
		//退出系统
		if !this.loop { 
			break
		}
	}
}
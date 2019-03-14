func Check_burte_force_ssh() {

	const SSHPORT =22
	const TIMEWAIT  ="TIME_WAIT"
	const SYNRECV  ="SYN_RECV"
	const CLOSING  = "CLOSING"
	const MAX_SSH_PORT_CONNECTION_NUM  =10

	var connectionnum=0
	nc,_:=net.Connections("all")//获取所有的连接状态
	for _,k:= range nc{
		if k.Laddr.Port== SSHPORT{
			if k.Status==TIMEWAIT||k.Status==SYNRECV||k.Status==CLOSING {//筛选出22端口的连接状态
				connectionnum+=1;//获取连接数目
			}
		}
	}
	if connectionnum>=MAX_SSH_PORT_CONNECTION_NUM { //如果超过设定的阀值，则报警。
		fmt.Print("有人在爆破我！！！！！")
	}

}

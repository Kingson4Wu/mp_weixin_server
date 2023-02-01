package proc

import (
	"github.com/kingson4wu/mp_weixin_server/common/bash"
	"log"
	"strconv"
	"strings"
)

/**

https://blog.csdn.net/aiyin5574/article/details/102123264

go 获得进程启动时间的两种方法

方法一：根据进行的pid，获得进程相关的文件夹，/proc/pid。将这个文件夹被修改的时间作为进程的启动时间。这样获取的进程启动时间不一定是准确的，但是相比方法二，花费的时间少。

var stat os.FileInfo
if stat, err = os.Lstat(fmt.Sprintf("/proc/%v", pid)); err != nil {
      return nil
}
proc.mtime = stat.ModTime().Unix()

方法二：根据进程相关的文件/proc/pid/stat中的内容获得进程启动时距离系统启动时的时间间隔，根据间隔计算出进程的启动时间。这种方法获得的时间更准确，但是性能可能比方法一差。

var (
    Uptime int64 // 系统启动时间戳
    scClkTck = int64(C.sysconf(C._SC_CLK_TCK))
)

func init() {
    sys := syscall.Sysinfo_t{}
    syscall.Sysinfo(&sys)
    Uptime = time.Now().Unix() - sys.Uptime
}

func ProcessStartTime(pid int) (ts time.Time) {
    buf, err := ioutil.ReadFile(fmt.Sprintf("/proc/%v/stat", pid))
    if err != nil {
        return time.Unix(0, 0)
    }
    if fields := strings.Fields(string(buf)); len(fields) > 22 {
        start, err := strconv.ParseInt(fields[21], 10, 0)
        if err == nil {
            if scClkTck > 0 {
                return time.Unix(Uptime+(start/scClkTck), 0)
            }
            return time.Unix(Uptime+(start/100), 0)
        }
    }
    return time.Unix(0, 0)
}
*/

/**
https://vimsky.com/examples/detail/golang-ex-github.com.shirou.gopsutil.process-Process---class.html
*/

/* func GetProcCmdline(procname string) (cmd []string, err error) {
	var proc *process.Process
	pid := getProcPID(procname)
	proc, err = process.NewProcess(int32(pid))
	cmd, err = proc.CmdlineSlice()
	return cmd, err
} */

func ExistProcName(name string) bool {

	result := bash.ExecShellCmd("ps -ef |grep " + name + " |grep -v 'grep'|wc -l")

	//去掉换行符
	count, err := strconv.Atoi(strings.TrimSpace(strings.Replace(result, "\n", "", -1)))
	if err != nil {
		log.Println("ExistProcName error :" + err.Error())
		return false
	}
	return count > 0
}

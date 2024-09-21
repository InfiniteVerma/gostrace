package main

type SysCallType int

const (
	read SysCallType = iota
	write
	open
	sys_close
	stat
	fstat
	lstat
	poll
	lseek
	mmap
	mprotect
	munmap
	brk
	rt_sigaction
	rt_sigprocmask
	rt_sigreturn
	ioctl
	pread64
	pwrite64
	readv
	writev
	access
	pipe
	sys_select
	sched_yield
	mremap
	msync
	mincore
	madvise
	shmget
	shmat
	shmctl
	dup
	dup2
	pause
	nanosleep
	getitimer
	alarm
	setitimer
	getpid
	sendfile
	socket
	connect
	accept
	sendto
	recvfrom
	sendmsg
	recvmsg
	shutdown
	bind
	listen
	getsockname
	getpeername
	socketpair
	setsockopt
	getsockopt
	clone
	fork
	vfork
	execve
	exit
	wait4
	kill
	uname
	semget
	semop
	semctl
	shmdt
	msgget
	msgsnd
	msgrcv
	msgctl
	fcntl
	flock
	fsync
	fdatasync
	truncate
	ftruncate
	getdents
	getcwd
	chdir
	fchdir
	rename
	mkdir
	rmdir
	creat
	link
	unlink
	symlink
	readlink
	chmod
	fchmod
	chown
	fchown
	lchown
	umask
	gettimeofday
	getrlimit
	getrusage
	sysinfo
	times
	ptrace
	getuid
	syslog
	getgid
	setuid
	setgid
	geteuid
	getegid
	setpgid
	getppid
	getpgrp
	setsid
	setreuid
	setregid
	getgroups
	setgroups
	setresuid
	getresuid
	setresgid
	getresgid
	getpgid
	setfsuid
	setfsgid
	getsid
	capget
	capset
	rt_sigpending
	rt_sigtimedwait
	rt_sigqueueinfo
	rt_sigsuspend
	sigaltstack
	utime
	mknod
	uselib
	personality
	ustat
	statfs
	fstatfs
	sysfs
	getpriority
	setpriority
	sched_setparam
	sched_getparam
	sched_setscheduler
	sched_getscheduler
	sched_get_priority_max
	sched_get_priority_min
	sched_rr_get_interval
	mlock
	munlock
	mlockall
	munlockall
	vhangup
	modify_ldt
	pivot_root
	_sysctl
	prctl
	arch_prctl
	adjtimex
	setrlimit
	chroot
	sync
	acct
	settimeofday
	mount
	umount2
	swapon
	swapoff
	reboot
	sethostname
	setdomainname
	iopl
	ioperm
	create_module
	init_module
	delete_module
	get_kernel_syms
	query_module
	quotactl
	nfsservctl
	getpmsg
	putpmsg
	afs_syscall
	tuxcall
	security
	gettid
	readahead
	setxattr
	lsetxattr
	fsetxattr
	getxattr
	lgetxattr
	fgetxattr
	listxattr
	llistxattr
	flistxattr
	removexattr
	lremovexattr
	fremovexattr
	tkill
	time
	futex
	sched_setaffinity
	sched_getaffinity
	set_thread_area
	io_setup
	io_destroy
	io_getevents
	io_submit
	io_cancel
	get_thread_area
	lookup_dcookie
	epoll_create
	epoll_ctl_old
	epoll_wait_old
	remap_file_pages
	getdents64
	set_tid_address
	restart_syscall
	semtimedop
	fadvise64
	timer_create
	timer_settime
	timer_gettime
	timer_getoverrun
	timer_delete
	clock_settime
	clock_gettime
	clock_getres
	clock_nanosleep
	exit_group
	epoll_wait
	epoll_ctl
	tgkill
	utimes
	vserver
	mbind
	set_mempolicy
	get_mempolicy
	mq_open
	mq_unlink
	mq_timedsend
	mq_timedreceive
	mq_notify
	mq_getsetattr
	kexec_load
	waitid
	add_key
	request_key
	keyctl
	ioprio_set
	ioprio_get
	inotify_init
	inotify_add_watch
	inotify_rm_watch
	migrate_pages
	openat
	mkdirat
	mknodat
	fchownat
	futimesat
	newfstatat
	unlinkat
	renameat
	linkat
	symlinkat
	readlinkat
	fchmodat
	faccessat
	pselect6
	ppoll
	unshare
	set_robust_list
	get_robust_list
	splice
	tee
	sync_file_range
	vmsplice
	move_pages
	utimensat
	epoll_pwait
	signalfd
	timerfd_create
	eventfd
	fallocate
	timerfd_settime
	timerfd_gettime
	accept4
	signalfd4
	eventfd2
	epoll_create1
	dup3
	pipe2
	inotify_init1
	preadv
	pwritev
	rt_tgsigqueueinfo
	perf_event_open
	recvmmsg
	fanotify_init
	fanotify_mark
	prlimit64
	name_to_handle_at
	open_by_handle_at
	clock_adjtime
	syncfs
	sendmmsg
	setns
	getcpu
	process_vm_readv
	process_vm_writev
	kcmp
	finit_module
	sched_setattr
	sched_getattr
	renameat2
	seccomp
	getrandom
	memfd_create
	kexec_file_load
	bpf
	execveat
	userfaultfd
	membarrier
	mlock2
	copy_file_range
	preadv2
	pwritev2
	pkey_mprotect
	pkey_alloc
	pkey_free
	statx
	io_pgetevents
	rseq
	pidfd_send_signal
	io_uring_setup
	io_uring_enter
	io_uring_register
	open_tree
	move_mount
	fsopen
	fsconfig
	fsmount
	fspick
	pidfd_open
	clone3
	close_range
	openat2
	pidfd_getfd
	faccessat2
	process_madvise
	epoll_pwait2
	mount_setattr
	quotactl_fd
	landlock_create_ruleset
	landlock_add_rule
	landlock_restrict_self
	memfd_secret
	process_mrelease
	futex_waitv
	set_mempolicy_home_node
	cachestat
	fchmodat2
	map_shadow_stack
	futex_wake
	futex_wait
	futex_requeue
	statmount
	listmount
	lsm_get_self_attr
	lsm_set_self_attr
	lsm_list_modules
	mseal
)

var sysCallNames = map[SysCallType]string{
	read:                    "read",
	write:                   "write",
	open:                    "open",
	sys_close:               "close",
	stat:                    "stat",
	fstat:                   "fstat",
	lstat:                   "lstat",
	poll:                    "poll",
	lseek:                   "lseek",
	mmap:                    "mmap",
	mprotect:                "mprotect",
	munmap:                  "munmap",
	brk:                     "brk",
	rt_sigaction:            "rt_sigaction",
	rt_sigprocmask:          "rt_sigprocmask",
	rt_sigreturn:            "rt_sigreturn",
	ioctl:                   "ioctl",
	pread64:                 "pread64",
	pwrite64:                "pwrite64",
	readv:                   "readv",
	writev:                  "writev",
	access:                  "access",
	pipe:                    "pipe",
	sys_select:              "select",
	sched_yield:             "sched_yield",
	mremap:                  "mremap",
	msync:                   "msync",
	mincore:                 "mincore",
	madvise:                 "madvise",
	shmget:                  "shmget",
	shmat:                   "shmat",
	shmctl:                  "shmctl",
	dup:                     "dup",
	dup2:                    "dup2",
	pause:                   "pause",
	nanosleep:               "nanosleep",
	getitimer:               "getitimer",
	alarm:                   "alarm",
	setitimer:               "setitimer",
	getpid:                  "getpid",
	sendfile:                "sendfile",
	socket:                  "socket",
	connect:                 "connect",
	accept:                  "accept",
	sendto:                  "sendto",
	recvfrom:                "recvfrom",
	sendmsg:                 "sendmsg",
	recvmsg:                 "recvmsg",
	shutdown:                "shutdown",
	bind:                    "bind",
	listen:                  "listen",
	getsockname:             "getsockname",
	getpeername:             "getpeername",
	socketpair:              "socketpair",
	setsockopt:              "setsockopt",
	getsockopt:              "getsockopt",
	clone:                   "clone",
	fork:                    "fork",
	vfork:                   "vfork",
	execve:                  "execve",
	exit:                    "exit",
	wait4:                   "wait4",
	kill:                    "kill",
	uname:                   "uname",
	semget:                  "semget",
	semop:                   "semop",
	semctl:                  "semctl",
	shmdt:                   "shmdt",
	msgget:                  "msgget",
	msgsnd:                  "msgsnd",
	msgrcv:                  "msgrcv",
	msgctl:                  "msgctl",
	fcntl:                   "fcntl",
	flock:                   "flock",
	fsync:                   "fsync",
	fdatasync:               "fdatasync",
	truncate:                "truncate",
	ftruncate:               "ftruncate",
	getdents:                "getdents",
	getcwd:                  "getcwd",
	chdir:                   "chdir",
	fchdir:                  "fchdir",
	rename:                  "rename",
	mkdir:                   "mkdir",
	rmdir:                   "rmdir",
	creat:                   "creat",
	link:                    "link",
	unlink:                  "unlink",
	symlink:                 "symlink",
	readlink:                "readlink",
	chmod:                   "chmod",
	fchmod:                  "fchmod",
	chown:                   "chown",
	fchown:                  "fchown",
	lchown:                  "lchown",
	umask:                   "umask",
	gettimeofday:            "gettimeofday",
	getrlimit:               "getrlimit",
	getrusage:               "getrusage",
	sysinfo:                 "sysinfo",
	times:                   "times",
	ptrace:                  "ptrace",
	getuid:                  "getuid",
	syslog:                  "syslog",
	getgid:                  "getgid",
	setuid:                  "setuid",
	setgid:                  "setgid",
	geteuid:                 "geteuid",
	getegid:                 "getegid",
	setpgid:                 "setpgid",
	getppid:                 "getppid",
	getpgrp:                 "getpgrp",
	setsid:                  "setsid",
	setreuid:                "setreuid",
	setregid:                "setregid",
	getgroups:               "getgroups",
	setgroups:               "setgroups",
	setresuid:               "setresuid",
	getresuid:               "getresuid",
	setresgid:               "setresgid",
	getresgid:               "getresgid",
	getpgid:                 "getpgid",
	setfsuid:                "setfsuid",
	setfsgid:                "setfsgid",
	getsid:                  "getsid",
	capget:                  "capget",
	capset:                  "capset",
	rt_sigpending:           "rt_sigpending",
	rt_sigtimedwait:         "rt_sigtimedwait",
	rt_sigqueueinfo:         "rt_sigqueueinfo",
	rt_sigsuspend:           "rt_sigsuspend",
	sigaltstack:             "sigaltstack",
	utime:                   "utime",
	mknod:                   "mknod",
	uselib:                  "uselib",
	personality:             "personality",
	ustat:                   "ustat",
	statfs:                  "statfs",
	fstatfs:                 "fstatfs",
	sysfs:                   "sysfs",
	getpriority:             "getpriority",
	setpriority:             "setpriority",
	sched_setparam:          "sched_setparam",
	sched_getparam:          "sched_getparam",
	sched_setscheduler:      "sched_setscheduler",
	sched_getscheduler:      "sched_getscheduler",
	sched_get_priority_max:  "sched_get_priority_max",
	sched_get_priority_min:  "sched_get_priority_min",
	sched_rr_get_interval:   "sched_rr_get_interval",
	mlock:                   "mlock",
	munlock:                 "munlock",
	mlockall:                "mlockall",
	munlockall:              "munlockall",
	vhangup:                 "vhangup",
	modify_ldt:              "modify_ldt",
	pivot_root:              "pivot_root",
	_sysctl:                 "_sysctl",
	prctl:                   "prctl",
	arch_prctl:              "arch_prctl",
	adjtimex:                "adjtimex",
	setrlimit:               "setrlimit",
	chroot:                  "chroot",
	sync:                    "sync",
	acct:                    "acct",
	settimeofday:            "settimeofday",
	mount:                   "mount",
	umount2:                 "umount2",
	swapon:                  "swapon",
	swapoff:                 "swapoff",
	reboot:                  "reboot",
	sethostname:             "sethostname",
	setdomainname:           "setdomainname",
	iopl:                    "iopl",
	ioperm:                  "ioperm",
	create_module:           "create_module",
	init_module:             "init_module",
	delete_module:           "delete_module",
	get_kernel_syms:         "get_kernel_syms",
	query_module:            "query_module",
	quotactl:                "quotactl",
	nfsservctl:              "nfsservctl",
	getpmsg:                 "getpmsg",
	putpmsg:                 "putpmsg",
	afs_syscall:             "afs_syscall",
	tuxcall:                 "tuxcall",
	security:                "security",
	gettid:                  "gettid",
	readahead:               "readahead",
	setxattr:                "setxattr",
	lsetxattr:               "lsetxattr",
	fsetxattr:               "fsetxattr",
	getxattr:                "getxattr",
	lgetxattr:               "lgetxattr",
	fgetxattr:               "fgetxattr",
	listxattr:               "listxattr",
	llistxattr:              "llistxattr",
	flistxattr:              "flistxattr",
	removexattr:             "removexattr",
	lremovexattr:            "lremovexattr",
	fremovexattr:            "fremovexattr",
	tkill:                   "tkill",
	time:                    "time",
	futex:                   "futex",
	sched_setaffinity:       "sched_setaffinity",
	sched_getaffinity:       "sched_getaffinity",
	set_thread_area:         "set_thread_area",
	io_setup:                "io_setup",
	io_destroy:              "io_destroy",
	io_getevents:            "io_getevents",
	io_submit:               "io_submit",
	io_cancel:               "io_cancel",
	get_thread_area:         "get_thread_area",
	lookup_dcookie:          "lookup_dcookie",
	epoll_create:            "epoll_create",
	epoll_ctl_old:           "epoll_ctl_old",
	epoll_wait_old:          "epoll_wait_old",
	remap_file_pages:        "remap_file_pages",
	getdents64:              "getdents64",
	set_tid_address:         "set_tid_address",
	restart_syscall:         "restart_syscall",
	semtimedop:              "semtimedop",
	fadvise64:               "fadvise64",
	timer_create:            "timer_create",
	timer_settime:           "timer_settime",
	timer_gettime:           "timer_gettime",
	timer_getoverrun:        "timer_getoverrun",
	timer_delete:            "timer_delete",
	clock_settime:           "clock_settime",
	clock_gettime:           "clock_gettime",
	clock_getres:            "clock_getres",
	clock_nanosleep:         "clock_nanosleep",
	exit_group:              "exit_group",
	epoll_wait:              "epoll_wait",
	epoll_ctl:               "epoll_ctl",
	tgkill:                  "tgkill",
	utimes:                  "utimes",
	vserver:                 "vserver",
	mbind:                   "mbind",
	set_mempolicy:           "set_mempolicy",
	get_mempolicy:           "get_mempolicy",
	mq_open:                 "mq_open",
	mq_unlink:               "mq_unlink",
	mq_timedsend:            "mq_timedsend",
	mq_timedreceive:         "mq_timedreceive",
	mq_notify:               "mq_notify",
	mq_getsetattr:           "mq_getsetattr",
	kexec_load:              "kexec_load",
	waitid:                  "waitid",
	add_key:                 "add_key",
	request_key:             "request_key",
	keyctl:                  "keyctl",
	ioprio_set:              "ioprio_set",
	ioprio_get:              "ioprio_get",
	inotify_init:            "inotify_init",
	inotify_add_watch:       "inotify_add_watch",
	inotify_rm_watch:        "inotify_rm_watch",
	migrate_pages:           "migrate_pages",
	openat:                  "openat",
	mkdirat:                 "mkdirat",
	mknodat:                 "mknodat",
	fchownat:                "fchownat",
	futimesat:               "futimesat",
	newfstatat:              "newfstatat",
	unlinkat:                "unlinkat",
	renameat:                "renameat",
	linkat:                  "linkat",
	symlinkat:               "symlinkat",
	readlinkat:              "readlinkat",
	fchmodat:                "fchmodat",
	faccessat:               "faccessat",
	pselect6:                "pselect6",
	ppoll:                   "ppoll",
	unshare:                 "unshare",
	set_robust_list:         "set_robust_list",
	get_robust_list:         "get_robust_list",
	splice:                  "splice",
	tee:                     "tee",
	sync_file_range:         "sync_file_range",
	vmsplice:                "vmsplice",
	move_pages:              "move_pages",
	utimensat:               "utimensat",
	epoll_pwait:             "epoll_pwait",
	signalfd:                "signalfd",
	timerfd_create:          "timerfd_create",
	eventfd:                 "eventfd",
	fallocate:               "fallocate",
	timerfd_settime:         "timerfd_settime",
	timerfd_gettime:         "timerfd_gettime",
	accept4:                 "accept4",
	signalfd4:               "signalfd4",
	eventfd2:                "eventfd2",
	epoll_create1:           "epoll_create1",
	dup3:                    "dup3",
	pipe2:                   "pipe2",
	inotify_init1:           "inotify_init1",
	preadv:                  "preadv",
	pwritev:                 "pwritev",
	rt_tgsigqueueinfo:       "rt_tgsigqueueinfo",
	perf_event_open:         "perf_event_open",
	recvmmsg:                "recvmmsg",
	fanotify_init:           "fanotify_init",
	fanotify_mark:           "fanotify_mark",
	prlimit64:               "prlimit64",
	name_to_handle_at:       "name_to_handle_at",
	open_by_handle_at:       "open_by_handle_at",
	clock_adjtime:           "clock_adjtime",
	syncfs:                  "syncfs",
	sendmmsg:                "sendmmsg",
	setns:                   "setns",
	getcpu:                  "getcpu",
	process_vm_readv:        "process_vm_readv",
	process_vm_writev:       "process_vm_writev",
	kcmp:                    "kcmp",
	finit_module:            "finit_module",
	sched_setattr:           "sched_setattr",
	sched_getattr:           "sched_getattr",
	renameat2:               "renameat2",
	seccomp:                 "seccomp",
	getrandom:               "getrandom",
	memfd_create:            "memfd_create",
	kexec_file_load:         "kexec_file_load",
	bpf:                     "bpf",
	execveat:                "execveat",
	userfaultfd:             "userfaultfd",
	membarrier:              "membarrier",
	mlock2:                  "mlock2",
	copy_file_range:         "copy_file_range",
	preadv2:                 "preadv2",
	pwritev2:                "pwritev2",
	pkey_mprotect:           "pkey_mprotect",
	pkey_alloc:              "pkey_alloc",
	pkey_free:               "pkey_free",
	statx:                   "statx",
	io_pgetevents:           "io_pgetevents",
	rseq:                    "rseq",
	pidfd_send_signal:       "pidfd_send_signal",
	io_uring_setup:          "io_uring_setup",
	io_uring_enter:          "io_uring_enter",
	io_uring_register:       "io_uring_register",
	open_tree:               "open_tree",
	move_mount:              "move_mount",
	fsopen:                  "fsopen",
	fsconfig:                "fsconfig",
	fsmount:                 "fsmount",
	fspick:                  "fspick",
	pidfd_open:              "pidfd_open",
	clone3:                  "clone3",
	close_range:             "close_range",
	openat2:                 "openat2",
	pidfd_getfd:             "pidfd_getfd",
	faccessat2:              "faccessat2",
	process_madvise:         "process_madvise",
	epoll_pwait2:            "epoll_pwait2",
	mount_setattr:           "mount_setattr",
	quotactl_fd:             "quotactl_fd",
	landlock_create_ruleset: "landlock_create_ruleset",
	landlock_add_rule:       "landlock_add_rule",
	landlock_restrict_self:  "landlock_restrict_self",
	memfd_secret:            "memfd_secret",
	process_mrelease:        "process_mrelease",
	futex_waitv:             "futex_waitv",
	set_mempolicy_home_node: "set_mempolicy_home_node",
	cachestat:               "cachestat",
	fchmodat2:               "fchmodat2",
	map_shadow_stack:        "map_shadow_stack",
	futex_wake:              "futex_wake",
	futex_wait:              "futex_wait",
	futex_requeue:           "futex_requeue",
	statmount:               "statmount",
	listmount:               "listmount",
	lsm_get_self_attr:       "lsm_get_self_attr",
	lsm_set_self_attr:       "lsm_set_self_attr",
	lsm_list_modules:        "lsm_list_modules",
	mseal:                   "mseal",
}

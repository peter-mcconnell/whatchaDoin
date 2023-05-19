#include <linux/bpf.h>
#include <bpf_helpers.h>

#define TASK_COMM_LEN 16

// this structure copied from "cat /sys/kernel/tracing/events/syscalls/sys_enter_execve/format"
struct syscalls_enter_execve {
    unsigned short common_type;
    unsigned char common_flags;
    unsigned char common_preempt_count;
    int common_pid;

    long int __syscall_nr;
    const char *filename;
    const char *const *argv;
    const char *const *envp;
};

SEC("tp/syscalls/sys_enter_execve")
int whatchadoin(struct syscalls_enter_execve *ctx)
{
	char fmt[] = "%s (pid: %d)\n";
  char comm[TASK_COMM_LEN];
  bpf_get_current_comm(&comm, sizeof(comm));
  bpf_trace_printk(fmt, sizeof(fmt), comm, bpf_get_current_pid_tgid() >> 32);

  return 0;
}

char _license[] SEC("license") = "GPL";

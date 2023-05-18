#include <linux/bpf.h>
#include <bpf_helpers.h>

int whatchadoin()
{
	return 0;
}

char _license[] SEC("license") = "GPL";

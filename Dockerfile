FROM scratch
COPY whatchadoin /whatchadoin
COPY bpf/whatchadoin_kern.o /bpf/whatchadoin_kern.o
CMD ["/whatchadoin"]

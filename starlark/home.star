load("starlark/lib/dag.star", "node", "edge", "is_runnable", "sequence")


def is_new_user(ctx, inputs):
    return ctx.cvt < 600

def main(args):
    recall = node("container_recall", {"type": "realtime"}, {"enabled": False}, None)
    rank = node("rank", None, {"enabled": True}, None)
    sequence(recall, rank)
    return 0

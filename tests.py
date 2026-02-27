import customlib_testing as t
import sys

show_missing = True
if "quiet" in sys.argv:
    t.show_success = False
    show_missing = False


t.g("/ws")
t.g("/metrics")
t.g("/session")

if __name__ == "__main__":
    t.list_missing(show_missing)
    t.run_tests()

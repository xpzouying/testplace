#!/usr/bin/env python
"""Provides NumberList and FrequencyDistribution, classes for statistics.

NumberList holds a sequence of numbers, and defines several statistical
operations (mean, stdev, etc.) FrequencyDistribution holds a mapping from
items (not necessarily numbers) to counts, and defines operations such as
Shannon entropy and frequency normalization.
"""

from math import sqrt, log, e
from random import choice, random
from Utils import indices

__author__ = "Rob Knight, Gavin Huttley, and Peter Maxwell"
__copyright__ = "Copyright 2007, The Cogent Project"
__credits__ = ["Rob Knight", "Peter Maxwell", "Gavin Huttley",
                    "Matthew Wakefield"]
__license__ = "GPL"
__version__ = "1.0.1"
__maintainer__ = "Rob Knight"
__email__ = "rob@spot.colorado.edu"
__status__ = "Production"

class NumberList(list):
    pass    #much code deleted
class FrequencyDistribution(dict):
    pass    #much code deleted

if __name__ == '__main__':    #code to execute if called from command-line
    pass    #do nothing - code deleted
#use this either for a simple example of how to use the module,
#or when the module can meaningfully be called as a script.

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

library LibSD {
    struct SD {
        int64 mean;
        int64 sqSum;
    }

    function add(SD storage self, int64 x_, int64 updatedCount_) internal {
        int64 oldMean_ = self.mean;
        // TODO: check (x_ - self.mean) underflow
        self.mean = self.mean + (x_ - self.mean) / updatedCount_;
        self.sqSum = self.sqSum + (x_ - self.mean) * (x_ - oldMean_);
    }

    function remove(SD storage self, int64 x_, int64 updatedCount_) internal {
        if (updatedCount_ == 0) {
            self.mean = 0;
            self.sqSum = 0;
            return;
        }

        int64 oldMean_ = self.mean;
        // TODO: check (x_ - self.mean) underflow
        self.mean = self.mean - (x_ - self.mean) / updatedCount_;
        self.sqSum = self.sqSum - (x_ - self.mean) * (x_ - oldMean_);
    }
}

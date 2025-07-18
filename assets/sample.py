# Python example for typing practice
import math
import random
from typing import List, Dict, Optional

class DataProcessor:
    """A class for processing numerical data."""
    
    def __init__(self, name: str):
        self.name = name
        self.data: List[float] = []
        self.processed = False
    
    def add_data(self, values: List[float]) -> None:
        """Add data points to the processor."""
        self.data.extend(values)
        self.processed = False
    
    def calculate_mean(self) -> float:
        """Calculate the arithmetic mean of the data."""
        if not self.data:
            return 0.0
        return sum(self.data) / len(self.data)
    
    def calculate_median(self) -> float:
        """Calculate the median of the data."""
        if not self.data:
            return 0.0
        
        sorted_data = sorted(self.data)
        n = len(sorted_data)
        
        if n % 2 == 0:
            return (sorted_data[n//2 - 1] + sorted_data[n//2]) / 2
        else:
            return sorted_data[n//2]
    
    def calculate_std_dev(self) -> float:
        """Calculate the standard deviation."""
        if len(self.data) < 2:
            return 0.0
        
        mean = self.calculate_mean()
        variance = sum((x - mean) ** 2 for x in self.data) / (len(self.data) - 1)
        return math.sqrt(variance)
    
    def find_outliers(self, threshold: float = 2.0) -> List[float]:
        """Find outliers using standard deviation method."""
        if len(self.data) < 3:
            return []
        
        mean = self.calculate_mean()
        std_dev = self.calculate_std_dev()
        
        outliers = []
        for value in self.data:
            z_score = abs(value - mean) / std_dev
            if z_score > threshold:
                outliers.append(value)
        
        return outliers

def main():
    """Main function demonstrating the data processor."""
    print("=== Data Processing Example ===")
    
    # Create processor
    processor = DataProcessor("Temperature Data")
    
    # Generate sample data
    sample_data = [23.5, 24.1, 22.8, 25.3, 23.9, 24.7, 23.2]
    processor.add_data(sample_data)
    
    # Calculate statistics
    mean = processor.calculate_mean()
    median = processor.calculate_median()
    std_dev = processor.calculate_std_dev()
    
    print(f"Dataset: {processor.name}")
    print(f"Mean: {mean:.2f}")
    print(f"Median: {median:.2f}")
    print(f"Standard deviation: {std_dev:.2f}")
    
    print("\nProcessing completed!")

if __name__ == "__main__":
    main()

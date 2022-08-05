# -*- coding: utf-8 -*-
"""
Created on Wed Jul 15 16:38:17 2020

@author: Icheng
"""

def random_sample_identities_forever(batch_size, num_samples_per_id, data_x,
                                     data_y, num_fa_images=0):
    """A generator that randomly selects a fixed number of entries per label.

    If false alarms are passed into this function, they should have a negative
    label, i.e., `data_y[i] < 0` if the i-th example corresponds to a false
    alarm.

    Parameters
    ----------
    batch_size : int
        The batch size.
    num_samples_per_id : int
        Number of examples per label in each batch. If the `batch_size` is not
        divisible by `num_samples_per_id` then the batch is filled with false
        alarms. A warning is printed if no false alarms are available to fill
        up the batch.
    data_x : List[string] | np.ndarray
        The data array; either a list of filenames or a tensor of input images.
    data_y : List[int] | np.ndarray
        The label array (either as list of one-dimensional numpy array).
    num_fa_images : Optional[int]
        Number of false alarm images to include in each batch; defaults to zero.

    Returns
    -------
    List[np.ndarray]
        Returns a list of length two where the first entry is the data array
        corresponding to `data_x` and the second entry is the label array
        corresponding to `data_y`. The elements in the list are of length
        `batch_size`.

    """
    assert (batch_size - num_fa_images) % num_samples_per_id == 0
    num_ids_per_batch = int((batch_size - num_fa_images) / num_samples_per_id)

    data_x = np.asarray(data_x)
    data_y = np.asarray(data_y)

    unique_y = np.unique(data_y[data_y >= 0])
    y_to_idx = {y: np.where(data_y == y)[0] for y in unique_y}
    fa_indices = np.where(data_y < 0)[0]

    while True:
        # Draw the desired number of identities.
        indices = np.random.choice(
            len(unique_y), num_ids_per_batch, replace=False)
        batch_unique_y = unique_y[indices]

        batch_x = np.zeros((batch_size, ) + data_x.shape[1:], data_x.dtype)
        batch_y = np.zeros((batch_size, ), data_y.dtype)
        e = 0
        for i, y in enumerate(batch_unique_y):
            num_samples = min(num_samples_per_id, len(y_to_idx[y]))
            indices = np.random.choice(y_to_idx[y], num_samples, replace=False)
            s, e = e, e + num_samples
            batch_x[s:e] = data_x[indices]
            batch_y[s:e] = y

        # Fill up remaining space with false alarms.
        num_samples = len(batch_x) - e
        if num_fa_images > 0:
            num_batch_fa_samples = min(num_samples, len(fa_indices))
            indices = np.random.choice(
                fa_indices, num_batch_fa_samples, replace=False)
            s, e = e, e + num_batch_fa_samples
            batch_x[s:e] = data_x[indices]
            batch_y[s:e] = data_y[indices]

        # If we need to add more data, random sample ids until we have reached
        # the batch size.
        num_samples = len(batch_x) - e
        num_tries = 0
        while num_samples > 0 and num_tries < 100:
            y = np.random.choice(unique_y)
            if y in batch_unique_y:
                # Find a target that we have not yet in this batch.
                num_tries += 1
                continue

            num_samples = min(num_samples, len(y_to_idx[y]))
            indices = np.random.choice(y_to_idx[y], num_samples, replace=False)
            s, e = e, e + num_samples
            batch_x[s:e] = data_x[indices]
            batch_y[s:e] = y
            num_samples = len(batch_x) - e

        if e < batch_size:
            print("ERROR: Failed to sample a full batch. Adding corrupt data.")
        yield [batch_x, batch_y]
        
        
dataset_dir = "MARS-evaluation-master"
dataset = Mars(dataset_dir, num_validation_y=0.1, seed=1234)
train_x, train_y, _ = dataset.read_train()
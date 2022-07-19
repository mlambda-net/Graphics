use ocl::ProQue;
use std::fmt::Error;

pub fn vector_add(a: Vec<i32>, b: Vec<i32>) -> Result<Vec<i32>, Error> {
    let src = r#"
        kernel void add(__global int* a, __global int* b, __global int *c) {
            int id = get_global_id(0);
            c[id] = a[id] + b[id];
        }
    "#;

    let hcl = ProQue::builder()
        .device(1)
        .src(src)
        .dims(a.len())
        .build()
        .unwrap();

    let a_buffer = hcl
        .buffer_builder()
        .copy_host_slice(a.as_slice())
        .build()
        .unwrap();

    let b_buffer = hcl
        .buffer_builder()
        .copy_host_slice(b.as_slice())
        .build()
        .unwrap();

    let c_buffer = hcl.create_buffer::<i32>().unwrap();

    let kernel = hcl
        .kernel_builder("add")
        .arg(&a_buffer)
        .arg(&b_buffer)
        .arg(&c_buffer)
        .build()
        .unwrap();

    unsafe { kernel.enq().unwrap() }

    let mut response = vec![0; a.len()];
    c_buffer.read(&mut response).enq().unwrap();

    return Ok(response);
}

pub fn vector_sum(vector: Vec<i32>) -> Result<i32, Error> {
    let src = r#"
        __kernel void sum(const __global int* vector, __local int* summation,  __global int *r) {
            
            int gid = get_global_id(0);
            int lid = get_local_id(0);
            int size = get_local_size(0);
            
            summation[lid] = vector[gid];
            
            
            for(int i = size / 2; i > 0; i /= 2) {
                barrier(CLK_LOCAL_MEM_FENCE);
            
                if (lid < i) {
                    summation[lid] += summation[lid + i];
                }
            }
            
            r[get_group_id(0)] = summation[0];
        }
    "#;

    let size = vector.len();

    let hcl = ProQue::builder()
        .device(1)
        .src(src)
        .dims(size)
        .build()
        .unwrap();

    let v_buffer = hcl
        .buffer_builder()
        .copy_host_slice(vector.as_slice())
        .build()
        .unwrap();

    let wgs = hcl.max_wg_size().unwrap();

    let len = (size / wgs) + 1;

    let r_buffer = ocl::Buffer::<i32>::builder()
        .queue(hcl.queue().clone())
        .len(len)
        .fill_val(Default::default())
        .build()
        .unwrap();

    let kernel = hcl
        .kernel_builder("sum")
        .arg(&v_buffer)
        .arg_local::<i32>(size)
        .arg(&r_buffer)
        .build()
        .unwrap();

    unsafe { kernel.enq().unwrap() }

    let mut response = vec![0; len];

    r_buffer.read(&mut response).enq().unwrap();

    return Ok(response.iter().sum());
}

pub fn vector_multiplication(a: Vec<i32>, b: Vec<i32>) -> Result<Vec<i32>, Error> {
    let src = r#"
        kernel void add(__global int* a, __global int* b, __global int *c) {
            int id = get_global_id(0);
            c[id] = a[id] * b[id];
        }
    "#;

    let hcl = ProQue::builder()
        .src(src)
        .device(1)
        .dims(a.len())
        .build()
        .unwrap();

    let a_buffer = hcl
        .buffer_builder()
        .copy_host_slice(a.as_slice())
        .build()
        .unwrap();

    let b_buffer = hcl
        .buffer_builder()
        .copy_host_slice(b.as_slice())
        .build()
        .unwrap();

    let c_buffer = hcl.create_buffer::<i32>().unwrap();

    let kernel = hcl
        .kernel_builder("add")
        .arg(&a_buffer)
        .arg(&b_buffer)
        .arg(&c_buffer)
        .build()
        .unwrap();

    unsafe { kernel.enq().unwrap() }

    let mut response = vec![0; a.len()];
    c_buffer.read(&mut response).enq().unwrap();

    return Ok(response);
}

pub fn vector_dot(a: Vec<i32>, b: Vec<i32>) -> Result<i32, Error> {
    let vector = vector_multiplication(a, b).unwrap();
    let dot = vector_sum(vector).unwrap();
    return Ok(dot);
}

#[cfg(test)]
pub mod test {
    use crate::math::{vector_add, vector_dot};
    use crate::math::{vector_multiplication, vector_sum};

    #[test]
    fn add_vector_test() {
        let a = vec![1; 65536];
        let b = vec![2; 65536];
        let c = vector_add(a, b).unwrap();

        assert_eq!(c, vec![3; 65536])
    }

    #[test]
    fn multiply_vector_test() {
        let a = vec![1; 2048];
        let b = vec![2; 2048];
        let c = vector_multiplication(a, b).unwrap();

        assert_eq!(c, vec![2; 2048])
    }

    #[test]
    fn summation_test() {
        let vector = [1; 16384];
        let r = vector_sum(vector.to_vec()).unwrap();
        assert_eq!(r, 16384)
    }

    #[test]
    fn vector_dot_test() {
        let vector = vec![1, 2, 3, 4, 5, 6, 7, 8];
        let r = vector_dot(vector.clone(), vector.clone()).unwrap();
        assert_eq!(204, r)
    }
}

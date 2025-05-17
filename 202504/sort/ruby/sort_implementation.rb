# ソートアルゴリズムの基本実装（クイックソート）
class SortImplementation
  def initialize
    # 初期化処理があれば行う
  end
  
  # reverse the array between lo and hi
  # lo: low, hi: high
  def reverse(arr, lo, hi)
    while lo < hi
      arr[lo], arr[hi] = arr[hi], arr[lo]
      lo += 1
      hi -= 1
    end
  end

  # make a copy from arr[lo] to arr[hi]
  def make_temp_array(arr, lo, hi)
    tmp_arr = arr[lo..hi]
    tmp_arr
  end

  # return the minimun length of a run from 23 ~ 64
  # so that arr.size/min is less than or equal to power of 2
  def merge_compute_minrun(n)
    run = 0
    while n >= 64
      run |= n & 1
      n >>= 1
    end
    n + run
  end

  def count_run(arr, s_run)
    increasing = true

    if s_run == arr.size - 1
      return [s_run, s_run, increasing, 1]
    else
      e_run = s_run

      if arr[s_run] > arr[s_run + 1]
        while arr[e_run] > arr[e_run + 1]
          e_run += 1
          if e_run == arr.size - 1
            break
          end
        end
        increasing = false
        return [s_run, e_run, increasing, e_run - s_run + 1]
      else
        while arr[e_run] <= arr[e_run + 1]
          e_run += 1
          if e_run == arr.size - 1
            break
          end
        end
        return [s_run, e_run, increasing, e_run - s_run + 1]
      end
    end
  end

  def bin_sort(arr, lo, hi, ext)
    (1..ext).each do |i|
      pos = 0
      st = lo
      ed = hi + i

      value = arr[ed]

      next if value >= arr[ed - 1]

      while st <= ed
        if st == ed
          if arr[st] > value
            pos = st
            break
          else
            pos = st + 1
            break
          end
        end
        mid = (st + ed) / 2
        if value >= arr[mid]
          st = mid + 1
        else
          ed = mid - 1
        end
      end

      if st > ed
        pos = st
      end

      (hi + i).downto(pos + 1) do |x|
        arr[x] = arr[x - 1]
      end
      arr[pos] = value
    end
  end

  def bisect_left(arr, val, lo, hi)
    while lo < hi
      mid = (lo + hi) / 2
      if arr[mid] < val
        lo = mid + 1
      else
        hi = mid
      end
    end
    lo
  end
  
  def bisect_right(arr, val, lo, hi)
    while lo < hi
      mid = (lo + hi) / 2
      if arr[mid] <= val
        lo = mid + 1
      else
        hi = mid
      end
    end
    lo
  end
  
  def gallop(arr, val, lo, hi, ltr)
    pos = ltr ? bisect_left(arr, val, lo, hi) : bisect_right(arr, val, lo, hi)
    pos
  end
  
  def merge(arr, stack, run)
    run_a = stack[run]
    run_b = stack[run + 1]
    
    new_run = [run_a[0], run_b[1], true, run_b[1] - run_a[0] + 1]
    stack[run] = new_run

    stack.delete_at(run + 1)

    if run_a[3] <= run_b[3]
      merge_low(arr, run_a, run_b, 7)
    else
      merge_high(arr, run_a, run_b, 7)
    end
  end

  def merge_low(arr, a, b, min_gallop)
    tmp_arr = make_temp_array(arr, a[0], a[1])

    k = a[0]
    i = 0
    j = b[0]

    gallop_thresh = min_gallop

    while true
      a_cnt = 0
      b_cnt = 0

      while i <= tmp_arr.size - 1 && j <= b[1]
        if tmp_arr[i] <= arr[j]
          arr[k] = tmp_arr[i]
          k += 1
          i += 1

          a_cnt += 1
          b_cnt = 0

          if i > tmp_arr.size - 1
            while j <= b[1]
              arr[k] = arr[j]
              k += 1
              j += 1
            end
            return
          end

          if a_cnt >= gallop_thresh
            break
          end
        else
          arr[k] = arr[j]
          k += 1
          j += 1

          a_cnt = 0
          b_cnt += 1

          if j > b[1]
            while i <= tmp_arr.size - 1
              arr[k] = tmp_arr[i]
              k += 1
              i += 1
            end
            return
          end

          if b_cnt >= gallop_thresh
            break
          end
        end
      end

      while true
        a_adv = gallop(tmp_arr, arr[j], i, tmp_arr.size, true)

        (i...a_adv).each do |x|
          arr[k] = tmp_arr[x]
          k += 1
        end

        a_cnt = a_adv - i
        i = a_adv

        if i > tmp_arr.size - 1
          while j <= b[1]
            arr[k] = arr[j]
            k += 1
            j += 1
          end
          return
        end

        arr[k] = arr[j]
        k += 1
        j += 1

        if j > b[1]
          while i < tmp_arr.size
            arr[k] = tmp_arr[i]
            k += 1
            i += 1
          end
          return
        end

        b_adv = gallop(arr, tmp_arr[i], j, b[1] + 1, true)
        (j...b_adv).each do |y|
          arr[k] = arr[y]
          k += 1
        end

        b_cnt = b_adv - j
        j = b_adv

        if j > b[1]
          while i <= tmp_arr.size - 1
            arr[k] = tmp_arr[i]
            k += 1
            i += 1
          end
          return
        end

        arr[k] = tmp_arr[i]
        i += 1
        k += 1

        if i > tmp_arr.size - 1
          while j <= b[1]
            arr[k] = arr[j]
            k += 1
            j += 1
          end
          return
        end

        if a_cnt < gallop_thresh && b_cnt < gallop_thresh
          break
        end
      end
      gallop_thresh += 1
    end
  end

  def merge_high(arr, a, b, min_gallop)
    tmp_arr = make_temp_array(arr, b[0], b[1])

    k = b[1]
    i = tmp_arr.size - 1
    j = a[1]

    gallop_thresh = min_gallop

    while true
      a_cnt = 0
      b_cnt = 0

      while i >= 0 && j >= a[0]
        if tmp_arr[i] >= arr[j]
          arr[k] = tmp_arr[i]
          k -= 1
          i -= 1

          a_cnt = 0
          b_cnt += 1

          if i < 0
            while j >= a[0]
              arr[k] = arr[j]
              k -= 1
              j -= 1
            end
            return
          end

          if b_cnt >= gallop_thresh
            break
          end
        else
          arr[k] = arr[j]
          k -= 1
          j -= 1

          a_cnt += 1
          b_cnt = 0

          if j < a[0]
            while i >= 0
              arr[k] = tmp_arr[i]
              k -= 1
              i -= 1
            end
            return
          end

          if a_cnt >= gallop_thresh
            break
          end
        end
      end

      while true
        a_adv = gallop(arr, tmp_arr[i], a[0], j + 1, false)

        j.downto(a_adv) do |x|
          arr[k] = arr[x]
          k -= 1
        end

        a_cnt = j - a_adv + 1
        j = a_adv - 1

        if j < a[0]
          while i >= 0
            arr[k] = tmp_arr[i]
            k -= 1
            i -= 1
          end
          return
        end

        arr[k] = tmp_arr[i]
        k -= 1
        i -= 1

        if i < 0
          while j >= a[0]
            arr[k] = arr[j]
            k -= 1
            j -= 1
          end
          return
        end

        b_adv = gallop(tmp_arr, arr[j], 0, i + 1, false)
        i.downto(b_adv) do |y|
          arr[k] = tmp_arr[y]
          k -= 1
        end

        b_cnt = i - b_adv + 1
        i = b_adv - 1

        if i < 0
          while j >= a[0]
            arr[k] = arr[j]
            k -= 1
            j -= 1
          end
          return
        end

        arr[k] = arr[j]
        j -= 1
        k -= 1

        if j < a[0]
          while i >= 0
            arr[k] = tmp_arr[i]
            k -= 1
            i -= 1
          end
          return
        end

        if a_cnt < gallop_thresh && b_cnt < gallop_thresh
          break
        end
      end
      gallop_thresh += 1
    end
  end

  def merge_collapse(arr, stack)
    while stack.size > 1
      if stack.size >= 3 && stack[-3][3] <= stack[-2][3] + stack[-1][3]
        if stack[-3][3] < stack[-1][3]
          merge(arr, stack, -3)
        else
          merge(arr, stack, -2)
        end
      elsif stack[-2][3] <= stack[-1][3]
        merge(arr, stack, -2)
      else
        break
      end
    end
  end

  def merge_force_collapse(arr, stack)
    while stack.size > 1
      merge(arr, stack, -2)
    end
  end

  def sort(arr)
    s = 0

    e = arr.size - 1
    stack = []
    min_run = merge_compute_minrun(arr.size)

    while s <= e
      run = count_run(arr, s)
      
      if run[2] == false
        reverse(arr, run[0], run[1])
        run[2] = true
      end

      if run[3] < min_run
        ext = [min_run - run[3], e - run[1]].min
        bin_sort(arr, run[0], run[1], ext)
        run[1] = run[1] + ext
        run[3] = run[3] + ext
      end

      stack.push(run)
      merge_collapse(arr, stack)
      s = run[1] + 1
    end
    merge_force_collapse(arr, stack)
    arr
  end
end
